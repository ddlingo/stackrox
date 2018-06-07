export const alerts = {
    countsByCluster: 'v1/alerts/summary/counts?group_by=CLUSTER*',
    countsByCategory: '/v1/alerts/summary/counts?group_by=CATEGORY*',
    alerts: '/v1/alerts?*',
    alertById: '/v1/alerts/*'
};

export const clusters = {
    list: 'v1/clusters',
    zip: 'api/extensions/clusters/zip'
};

export const benchmarks = {
    configs: 'v1/benchmarks/configs',
    benchmarkScans: 'v1/benchmarks/scans?benchmarkId=*',
    scanHostResults: '/v1/benchmarks/scans/*/*',
    scans: '/v1/benchmarks/scans/*',
    triggers: 'v1/benchmarks/triggers/*',
    summary: 'v1/benchmarks/summary/scans*'
};

export const risks = {
    riskyDeployments: 'v1/deployments*'
};

export const search = {
    globalSearchWithResults: '/v1/search?query=Stale:false+Cluster:remote',
    globalSearchWithNoResults: '/v1/search?query=Stale:false+Cluster:',
    options: '/v1/search/metadata/options*'
};

export const images = {
    list: '/v1/images*'
};

export const auth = {
    authProviders: 'v1/authProviders*',
    authStatus: '/v1/auth/status'
};
