/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { Typography, Table, Card } from 'antd';
import { useEffect, useMemo, useState } from 'react';
// https://charts.ant.design/en/manual/case
import { Column, Pie } from '@ant-design/plots';
import { getAllIncomes } from '../../../api/economics/incomes/getAll';
import { ZoomInOutlined } from '@ant-design/icons';
import { getAllIncomesByProjectId } from '../../../api/economics/incomes/getAllByProjectId';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import IncomeStatus from '../../status/IncomeStatus';
import { Income } from '../../../interfaces/income';
import { useGetExternalCompanies, useGetLoggedInUserId } from '../../../hooks';
import { formatNumberWithSpaces } from '../../../helpers/stringFormatting';

const { Text, Title, Link } = Typography;

const calculateTotalAmountAndTax = (incomes: Income[], getVendorName : (id : string) => string) => {
    let totalAmount = 0;
    let totalTax = 0;
    let totalIncomes = 0;
    const columnGraphData = []
    const pieGraphData = []
    const pieGraphTaxData = []
  
    for (const income of incomes) {
        columnGraphData.push({
            vendor: getVendorName(income.vendor),
            amount: income.amount,
            tax: income.tax,
            income_category: income.income_category,
            project_id: income.project_id
        });
        pieGraphData.push({
            income_category: income.income_category,
            amount: income.amount,
        })
        pieGraphTaxData.push({
            income_category: income.income_category,
            tax: income.tax,
        })
        totalIncomes += 1
        totalAmount += income.amount;
        totalTax += income.tax;
    }
  
    return { pieGraphTaxData, pieGraphData, columnGraphData, totalIncomes, totalAmount, totalTax };
  }

const incomesTabList = [
    {
      key: 'incomes',
      label: 'All incomes',
    },
    {
      key: 'summary',
      label: 'Income summary',
    },
  ];

const economicsColumns = [
    {
        title: 'Vendor',
        dataIndex: 'vendor',
        key: 'vendor'
    },
    {
        title: 'Description',
        dataIndex: 'description',
        key: 'description'
    },
    {
        title: 'Cost',
        dataIndex: 'cost',
        key: 'cost'
    },
    {
      title: 'Tax',
      dataIndex: 'tax',
      key: 'tax'
    },
    {
      title: 'Status',
      dataIndex: 'status',
      key: 'status'
    },
    {
        title: 'Income date',
        dataIndex: 'income_date',
        key: 'income_date'
    },
    {
      title: '',
      dataIndex: 'operations',
      key: 'operations'
    },
  ];

const Incomes = ({ project } : { project: string }) => {
    const loggedInUserId = useGetLoggedInUserId();
    const externalCompanies = useGetExternalCompanies();
    const [activeTab, setActiveTab] = useState<string>('incomes')
    const [incomes, setIncomes] = useState<Array<any>>([]);

    const getVendorName = (id : string) => externalCompanies?.[id]?.company_name;

    useEffect(() => {
        if (loggedInUserId && project === 'all') {
            getAllIncomes(loggedInUserId).then(response => {
                setIncomes(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        } else if (loggedInUserId) {
            getAllIncomesByProjectId(loggedInUserId, project).then(response => {
                setIncomes(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUserId, project]);

      const onHandleChangeActiveTab = (tab : string) => setActiveTab(tab);

      const economicsData: Array<any> = useMemo(() => {
        const incomesListItem = incomes && incomes.map((income : Income) => {
        return {                    
            vendor: <Link href={`/external-company/${income.vendor}`}>{getVendorName(income.vendor)}</Link>,
            description: <Text>{income.description}</Text>,
            cost: <Text>{formatNumberWithSpaces(income.amount)} {income.currency}</Text>,
            tax: <Text>{formatNumberWithSpaces(income.tax)} {income.currency}</Text>,
            status: <IncomeStatus status={income.status}/>,
            income_date: <Text>{formatDateTimeToYYYYMMDDHHMM(income.income_date)}</Text>,
            operations: <Link href={`/income/${income.id}`}><ZoomInOutlined /></Link>
          }
        })
        return incomesListItem;
    }, [project, incomes]);

    if (!economicsData) return null;

    const {pieGraphTaxData, pieGraphData, columnGraphData, totalIncomes, totalAmount, totalTax} = calculateTotalAmountAndTax(incomes, getVendorName)

    const columnShartConfig = {
      data: columnGraphData,
      xField: 'income_category',
      yField: 'amount',
      isStack: true,
      isGroup: true,
      groupField: 'income_category',
      seriesField: 'vendor',
    };

    const pieShartConfig = {
      appendPadding: 40,
      data: pieGraphData,
      angleField: 'amount',
      colorField: 'income_category',
    };

    const pieShartTaxConfig = {
      appendPadding: 40,
      data: pieGraphTaxData,
      angleField: 'tax',
      colorField: 'income_category',
    };


    const incomesContentList: Record<string, React.ReactNode> = {
      incomes:  <Table size="small" columns={economicsColumns} dataSource={economicsData} />,
      summary: (
          <div>
              <div style={{padding: '24px 24px 16px 24px', display: 'flex'}}>
                  <Text style={{paddingRight: '8px'}} strong>Total amount:</Text><Text  style={{paddingRight: '20px'}} >{formatNumberWithSpaces(totalAmount)} SEK</Text>
                  <Text style={{paddingRight: '8px'}} strong>Total tax:</Text><Text style={{paddingRight: '20px'}}>{formatNumberWithSpaces(totalTax)} SEK</Text>
                  <Text style={{paddingRight: '8px'}} strong>Total incomes:</Text><Text style={{paddingRight: '20px'}}>{formatNumberWithSpaces(totalIncomes)}</Text>
              </div>
              <div style={{padding: '16px'}}>
                  <Column {...columnShartConfig} />
                  <div style={{display: 'flex'}}>
                      <div style={{width: '49%', marginRight: '1%', marginTop: '48px'}}>
                      <Title style={{textAlign: 'center', marginTop: '24px'}} level={2}>Costs</Title>
                      <Pie {...pieShartConfig} />
                      </div>
                      <div style={{width: '59%', marginLeft: '1%', marginTop: '48px'}}>
                      <Title style={{textAlign: 'center', marginTop: '24px'}} level={2}>Taxes</Title>
                      <Pie {...pieShartTaxConfig} />
                      </div>
                  </div>
              </div>
          </div>
      )
    }

    return  (
        <Card 
            style={{ height: 'fit-content', padding: 0}}
            tabList={incomesTabList}
            activeTabKey={activeTab}
            bodyStyle={{padding: '0px'}}
            onTabChange={onHandleChangeActiveTab}
            >
            {incomesContentList[activeTab]}
        </Card>
    );

}

export default Incomes;