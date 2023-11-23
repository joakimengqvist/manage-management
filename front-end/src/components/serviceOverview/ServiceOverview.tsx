/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
// https://charts.ant.design/en/manual/case
import { FlowAnalysisGraph } from '@ant-design/graphs';


const Services: React.FC = () => {

    const data = {
        nodes: [
            {
                id: '0',
                value: {
                    title: 'Mongo DB',
                    items: [
                    {
                        text: 'Log source',
                    },
                    ],
                },
            },
            {
                id: '10',
                value: {
                title: 'Postgress',
                items: [
                    {
                        text: 'Database source',
                    },
                ],
                },
            },
            {
                id: '100',
                value: {
                title: 'Broker-service',
                items: [
                    {
                        text: 'API',
                    },
                ],
                },
            },

            {
                id: '1001',
                value: {
                title: 'Logger-service',
                items: [
                    {
                        text: 'Logs',
                    },
                ],
                },
            },
            {
                id: '1002',
                value: {
                title: 'Authentication-service',
                items: [
                    {
                        text: 'Authentication | users',
                    },
                ],
                },
            },
            {
                id: '1003',
                value: {
                title: 'Economics-service',
                items: [
                    {
                        text: 'incomes | expenses',
                    },
                ],
                },
            },
            {
                id: '1004',
                value: {
                title: 'Mail-service',
                items: [
                    {
                        text: 'email traffic',
                    },
                ],
                },
            },
            {
                id: '1005',
                value: {
                title: 'Project-service',
                items: [
                    {
                        text: 'projects',
                    },
                ],
                },
            },
            {
                id: '1006',
                value: {
                title: 'Notes-service',
                items: [
                    {
                        text: 'notes',
                    },
                ],
                },
            },
            {
                id: '1007',
                value: {
                title: 'External-company-service',
                items: [
                    {
                        text: 'External companies',
                    },
                ],
                },
            },
            {
                id: '1008',
                value: {
                title: 'Product-service',
                items: [
                    {
                        text: 'Products',
                    },
                ],
                },
            },
            {
                id: '1009',
                value: {
                title: 'Invoice-service',
                items: [
                    {
                        text: 'Invoices',
                    },
                ],
                },
            },
            {
                id: '10000',
                value: {
                title: 'Front-end',
                items: [
                    {
                        text: 'React | Vite',
                    },
                ],
                },
            },
        ],
        edges: [
            {
                source: '0',
                target: '1001',
            },
            {
                source: '10',
                target: '1002',
            },
            {
                source: '10',
                target: '1003',
            },
            {
                source: '10',
                target: '1004',
            },
            {
                source: '10',
                target: '1005',
            },
            {
                source: '10',
                target: '1006',
            },
            {
                source: '10',
                target: '1007',
            },
            {
                source: '10',
                target: '1008',
            },
            {
                source: '10',
                target: '1009',
            },
            {
                source: '1001',
                target: '1002',
            },
            {
                source: '1001',
                target: '1003',
            },

            {
                source: '1001',
                target: '1004',
            },
            {
                source: '1001',
                target: '1005',
            },
            {
                source: '1001',
                target: '1006',
            },
            {
                source: '1001',
                target: '1007',
            },
            {
                source: '1001',
                target: '1008',
            },
            {
                source: '1001',
                target: '1009',
            },
            {
                source: '1002',
                target: '100',
            },
            {
                source: '1003',
                target: '100',
            },
            {
                source: '1004',
                target: '100',
            },
            {
                source: '1005',
                target: '100',
            },
            {
                source: '1006',
                target: '100',
            },
            {
                source: '1007',
                target: '100',
            },
            {
                source: '1008',
                target: '100',
            },
            {
                source: '1009',
                target: '100',
            },
            {
                source: '100',
                target: '10000',
            },


        ],
      };
      const config = {
        data,
        layout: {
          rankdir: 'TB',
          ranksepFunc: () => 8,
        },
        nodeCfg: {
          anchorPoints: [
            [0.5, 0],
            [0.5, 1],
          ],
        },
        edgeCfg: {
          type: 'polyline',
        },
        markerCfg: (cfg: { id: string; }) => {
          return {
            position: 'bottom',
            show: data.edges.filter((item) => item.source === cfg.id)?.length,
          };
        },
        behaviors: ['drag-node'],
      };

    return <FlowAnalysisGraph {...config}/>;

}

export default Services;