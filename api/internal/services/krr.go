package services

import (
	"context"
	"fmt"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log/slog"
	"strings"
	"time"
)

type KRRService struct {
	log *slog.Logger
}

func NewKRRService(log *slog.Logger) *KRRService {
	return &KRRService{
		log: log,
	}
}

func (s *KRRService) Recs(ctx context.Context) (string, error) {
	const op = "services.krr.Recs"

	log := s.log.With(slog.String("op", op))

	log.Info("creating k8s client")

	client, err := createInClusterClient()
	if err != nil {
		log.Error("failed to create k8s client", slog.Any("error", err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("k8s client created, creating krr pod")

	pod, err := createPod(ctx, client, "default")
	if err != nil {
		log.Error("failed to create krr pod", slog.Any("error", err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("krr pod created, waiting for pod to finish")
	// todo repository for pod crud
	// todo return status (error or finish)
	err = waitForPodCompletion(ctx, client, "default", pod.Name)
	if err != nil {
		log.Error("pod failed", slog.Any("error", err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("pod finished, collecting logs")

	logs, err := getPodLogs(ctx, client, "default", pod.Name)
	if err != nil {
		log.Error("failed to collect logs", slog.Any("error", err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Info("logs collected")

	client.CoreV1().Pods("default").Delete(context.Background(), pod.Name, metav1.DeleteOptions{})

	return logs, nil
}

func createInClusterClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func createPod(
	ctx context.Context,
	clientset *kubernetes.Clientset,
	namespace string,
) (*v1.Pod, error) {
	//todo: move params to config
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "example-pod",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    "example-container",
					Image:   "us-central1-docker.pkg.dev/genuine-flight-317411/devel/krr:v1.8.3",
					Command: []string{"sh", "-c", "python krr.py simple --prometheus-url=http://prometheus-server:80 -q -f json"},
				},
			},
			RestartPolicy: v1.RestartPolicyNever,
		},
	}

	return clientset.CoreV1().Pods(namespace).Create(ctx, pod, metav1.CreateOptions{})
}

func waitForPodCompletion(ctx context.Context, clientset *kubernetes.Clientset, namespace, podName string) error {
	for {
		pod, err := clientset.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
		if err != nil {
			return err
		}

		if pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
			return nil
		}
		time.Sleep(2 * time.Second)
	}
}

func getPodLogs(ctx context.Context, clientset *kubernetes.Clientset, namespace, podName string) (string, error) {
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &v1.PodLogOptions{})
	logs, err := req.Stream(ctx)
	if err != nil {
		return "", err
	}
	defer logs.Close()

	var result strings.Builder
	buf := make([]byte, 2000)
	for {
		n, err := logs.Read(buf)
		if n > 0 {
			result.Write(buf[:n])
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return result.String(), nil
}
