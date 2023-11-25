/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { Typography, Table, Card } from 'antd';
import { useEffect, useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../../types/state';
// https://charts.ant.design/en/manual/case
import { Column, Pie } from '@ant-design/plots';
import { ZoomInOutlined } from '@ant-design/icons';
import { getAllExpenses } from '../../../api/economics/expenses/getAll';
import { getAllExpensesByProjectId } from '../../../api/economics/expenses/getAllByProjectId';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import ExpenseStatus from '../../status/ExpenseStatus';
import { ExpenseObject } from '../../../types/expense';

const { Text, Title, Link } = Typography;

const calculateTotalAmountAndTax = (expenses: ExpenseObject[], getVendorName : (id: string) => string | undefined) => {
    let totalAmount = 0;
    let totalTax = 0;
    let totalExpenses = 0;
    const columnGraphData = []
    const pieGraphData = []
    const pieGraphTaxData = []
  
    for (const expense of expenses) {
        columnGraphData.push({
            vendor: getVendorName(expense.vendor),
            amount: expense.amount,
            tax: expense.tax,
            expense_category: expense.expense_category,
            project_id: expense.project_id
        });
        pieGraphData.push({
            expense_category: expense.expense_category,
            amount: expense.amount,
        })
        pieGraphTaxData.push({
            expense_category: expense.expense_category,
            tax: expense.tax,
        })
        totalExpenses += 1
        totalAmount += expense.amount;
        totalTax += expense.tax;
    }
  
    return { pieGraphTaxData, pieGraphData, columnGraphData, totalExpenses, totalAmount, totalTax };
  }

const expensesTabList = [
    {
      key: 'expenses',
      label: 'All expenses',
    },
    {
      key: 'summary',
      label: 'Expense summary',
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
        title: 'Expense date',
        dataIndex: 'expense_date',
        key: 'expense_date'
    },
    {
      title: '',
      dataIndex: 'operations',
      key: 'operations'
    },
  ];

const Expenses = ({ project } : { project: string }) => {
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const externalCompanies = useSelector((state : State) => state.application.externalCompanies);
    const [activeTab, setActiveTab] = useState<string>('expenses')
    const [expenses, setExpenses] = useState<Array<any>>([]);

    const getVendorName = (id : string) => externalCompanies.find(company => company.id === id)?.company_name;

    useEffect(() => {
        if (loggedInUserId && project === 'all') {
            getAllExpenses(loggedInUserId).then(response => {
                setExpenses(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        } else if (loggedInUserId) {
            getAllExpensesByProjectId(loggedInUserId, project).then(response => {
                setExpenses(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUserId, project]);

      

      const onHandleChangeActiveTab = (tab : string) => setActiveTab(tab);

      const economicsData: Array<any> = useMemo(() => {
        const expensesListItem = expenses && expenses.map((expense : ExpenseObject) => {
        return {                    
            vendor: <Link href={`/external-company/${expense.vendor}`}>{getVendorName(expense.vendor)}</Link>,
            description: <Text>{expense.description}</Text>,
            cost: <Text>{expense.amount} {expense.currency}</Text>,
            tax: <Text>{expense.tax} {expense.currency}</Text>,
            payment_method: <Text>{expense.payment_method}</Text>,
            status: <ExpenseStatus status={expense.status}/>,
            expense_date: <Text>{formatDateTimeToYYYYMMDDHHMM(expense.expense_date)}</Text>,
            operations: <Link href={`/expense/${expense.id}`}><ZoomInOutlined /></Link>
          }
        })
        return expensesListItem;
    }, [project, expenses])

    if (!economicsData) return null;

    const {pieGraphTaxData, pieGraphData, columnGraphData, totalExpenses, totalAmount, totalTax} = calculateTotalAmountAndTax(expenses, getVendorName)

    const columnShartConfig = {
      data: columnGraphData,
      xField: 'expense_category',
      yField: 'amount',
      isStack: true,
      isGroup: true,
      groupField: 'expense_category',
      seriesField: 'vendor',
    };

    const pieShartConfig = {
      appendPadding: 40,
      data: pieGraphData,
      angleField: 'amount',
      colorField: 'expense_category',
    };

    const pieShartTaxConfig = {
      appendPadding: 40,
      data: pieGraphTaxData,
      angleField: 'tax',
      colorField: 'expense_category',
    };


    const expensesContentList: Record<string, React.ReactNode> = {
      expenses:  <Table size="small" columns={economicsColumns} dataSource={economicsData} />,
      summary: (
          <div>
              <div style={{padding: '24px 24px 16px 24px', display: 'flex'}}>
                  <Text style={{paddingRight: '8px'}} strong>Total amount:</Text><Text  style={{paddingRight: '20px'}} >{totalAmount}</Text>
                  <Text style={{paddingRight: '8px'}} strong>Total tax:</Text><Text style={{paddingRight: '20px'}}>{totalTax}</Text>
                  <Text style={{paddingRight: '8px'}} strong>Total expenses:</Text><Text style={{paddingRight: '20px'}}>{totalExpenses}</Text>
              </div>
              <div style={{padding: '16px'}}>
                  <Column {...columnShartConfig} />
                  <div style={{display: 'flex'}}>
                      <div style={{width: '59%', marginRight: '1%', marginTop: '48px'}}>
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
            tabList={expensesTabList}
            activeTabKey={activeTab}
            bodyStyle={{padding: '0px'}}
            onTabChange={onHandleChangeActiveTab}
            >
            {expensesContentList[activeTab]}
        </Card>
    );

}

export default Expenses;