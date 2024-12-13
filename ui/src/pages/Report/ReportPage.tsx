/* eslint-disable complexity */
/* eslint-disable import/no-extraneous-dependencies */

import './ReportPage.scss';
import {Button, Skeleton} from '@gravity-ui/uikit';
import ReportTable from '../../components/ReportTable/ReportTable';
import {useState} from 'react';
import KRRReport from '../../types/KRRReport';
import {defaultReport} from '../../types';
import axios from 'axios';
import recommendationsCfg from '../../config/recommendations';

const ReportPage = () => {
    const [reportData, setReportData] = useState<KRRReport>(defaultReport);
    const [reportDataLoading, setReportDataLoading] = useState<boolean>(false);

    const getReportData = () => {
        setReportDataLoading(true);
        // todo: move to service
        console.log(recommendationsCfg);
        axios.get(`${recommendationsCfg.url}/recommend`).then((res) => {
            setReportData(res.data);
            setReportDataLoading(false);
        });
    };

    const showTable = reportData.scans.length > 0 && !reportDataLoading;
    const showNoDataMessage = reportData.scans.length === 0 && !reportDataLoading;

    const tableRows = reportData.scans.map((e, i) => ({
        number: i + 1,
        namespace: e.object.namespace,
        name: e.object.name,
        pods: e.object.pods.reduce((acc, p) => acc + (p.deleted ? 0 : 1), 0),
        old_pods: e.object.pods.reduce((acc, p) => acc + (p.deleted ? 1 : 0), 0),
        type: e.object.kind,
        container: e.object.container,
        cpu_diff: 'hz',
        cpu_requests: `${e.object.allocations.requests.cpu} -> ${e.recommended.requests.cpu.value}`,
        cpu_limits: `${e.object.allocations.limits.cpu} -> ${e.recommended.limits.cpu.value}`,
        memory_diff: 'hz',
        memory_requests: `${e.object.allocations.requests.memory} -> ${e.recommended.requests.memory.value}`,
        memory_limits: `${e.object.allocations.limits.memory} -> ${e.recommended.limits.memory.value}`,
    }));

    return (
        <div className={'page'}>
            <h1>KRR Reports</h1>
            <Button className={'report-button'} onClick={getReportData}>
                Run krr
            </Button>
            {reportDataLoading && <Skeleton />}
            {showTable && <ReportTable data={tableRows} />}
            {showNoDataMessage && <p>No data</p>}
        </div>
    );
};

export default ReportPage;
