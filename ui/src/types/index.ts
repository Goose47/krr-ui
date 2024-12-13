import KRRReport from './KRRReport';

export * from './KRRReport';
export * from './KRRReportEntry';

const defaultReport: KRRReport = {
    scans: [],
    score: 0,
    description: '',
    errors: [],
};

export {defaultReport};
