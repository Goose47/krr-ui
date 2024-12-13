import KRRReportEntry from './KRRReportEntry';

interface KRRReport {
    scans: KRRReportEntry[];
    score: number;
    description: string;
    errors: any[];
}

export default KRRReport;
