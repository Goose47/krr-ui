import './ReportTable.scss';
import {Table} from '@gravity-ui/uikit';
import {FC} from 'react';

const columns = [
    {id: 'number', name: 'Number'},
    {id: 'namespace', name: 'Namespace'},
    {id: 'name', name: 'Name'},
    {id: 'pods', name: 'Pods'},
    {id: 'old_pods', name: 'Old Pods'},
    {id: 'type', name: 'Type'},
    {id: 'container', name: 'Container'},
    {id: 'cpu_diff', name: 'CPU Diff'},
    {id: 'cpu_requests', name: 'CPU Requests'},
    {id: 'cpu_limits', name: 'CPU Limits'},
    {id: 'memory_diff', name: 'Memory Diff'},
    {id: 'memory_requests', name: 'Memory Requests'},
    {id: 'memory_limits', name: 'Memory Limits'},
];

interface ComponentProps {
    data: any[];
}

const ReportTable: FC<ComponentProps> = ({data}) => {
    return <Table className={'report-table'} data={data} columns={columns} />;
};

export default ReportTable;
