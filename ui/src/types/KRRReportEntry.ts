interface KRRReportEntry {
    object: {
        name: string;
        namespace: string;
        container: string;
        kind: string;
        pods: {
            name: string;
            deleted: boolean;
        }[];
        allocations: {
            requests: {
                cpu: number | string;
                memory: number | string;
            };
            limits: {
                cpu: number | string;
                memory: number | string;
            };
            info: any;
        };
        warnings: string[];
    };
    recommended: {
        requests: {
            cpu: {
                value: number | string;
                severity: string;
            };
            memory: {
                value: number | string;
                severity: number | string;
            };
        };
        limits: {
            cpu: {
                value: number | string;
                severity: number | string;
            };
            memory: {
                value: number | string;
                severity: number | string;
            };
        };
        info: {
            cpu: number | string;
            memory: number | string;
        };
    };
    severity: string;
}

export default KRRReportEntry;
