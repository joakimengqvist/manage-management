/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { Typography, Button, Table, Card } from 'antd';
import { useEffect, useMemo, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../../types/state';
import { cardShadow } from '../../../enums/styles';
// https://charts.ant.design/en/manual/case
import { Column, Pie } from '@ant-design/plots';
import { getAllProjectExpenses } from '../../../api/economics/expenses/getAllProjectExpenses';
import { useNavigate } from 'react-router-dom';
import { getAllProjectExpensesByProjectId } from '../../../api/economics/expenses/getAllProjectExpensesByProjectId';
import { ExpenseAndIncomeStatus, PaymentStatusTypes } from '../../tags/ExpenseAndIncomeStatus';

const { Text, Title } = Typography;

const calculateTotalAmountAndTax = (expenses: ExpenseObject[]) => {
    let totalAmount = 0;
    let totalTax = 0;
    let totalExpenses = 0;
    const columnGraphData = []
    const pieGraphData = []
    const pieGraphTaxData = []
  
    for (const expense of expenses) {
        columnGraphData.push({
            vendor: expense.vendor,
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

type ExpenseObject = {
	expense_id: string,
	project_id: string,
  expense_date: any,
	expense_category: string,
	vendor: string,
	description: string,
	amount: number,
	tax: number,
  status: PaymentStatusTypes,
	currency: string,
	payment_method: string,
	created_by: string,
	created_at: any,
	modified_by: any,
	modified_at: any
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
        title: 'Payment method',
        dataIndex: 'payment_method',
        key: 'payment_method'

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
    const navigate = useNavigate();
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const externalCompanies = useSelector((state : State) => state.application.externalCompanies);
    const [activeTab, setActiveTab] = useState<string>('expenses')
    const [expenses, setExpenses] = useState<Array<any>>([]);

    useEffect(() => {
        if (loggedInUserId && project === 'all') {
            getAllProjectExpenses(loggedInUserId).then(response => {
                setExpenses(response)
            }).catch(error => {
                console.log('error fetching', error)
            })
        } else if (loggedInUserId) {
            getAllProjectExpensesByProjectId(loggedInUserId, project).then(response => {
                setExpenses(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUserId, project]);

      const {pieGraphTaxData, pieGraphData, columnGraphData, totalExpenses, totalAmount, totalTax} = calculateTotalAmountAndTax(expenses)

      const onHandleChangeActiveTab = (tab : string) => setActiveTab(tab);

      const getVendorName = (id : string) => externalCompanies.find(company => company.id === id)?.name;
      

      const economicsData: Array<any> = useMemo(() => {
        const expensesListItem = expenses.map((expense : ExpenseObject) => {
        return {                    
            vendor: <Button type="link" onClick={() => navigate(`/external-company/${expense.vendor}`)}>{getVendorName(expense.vendor)}</Button>,
            description: <Text>{expense.description}</Text>,
            cost: <Text>{expense.amount} {expense.currency}</Text>,
            tax: <Text>{expense.tax} {expense.currency}</Text>,
            payment_method: <Text>{expense.payment_method}</Text>,
            status: <ExpenseAndIncomeStatus status={expense.status}/>,
            expense_date: <Text>{expense.expense_date}</Text>,
            operations: <Button type="link" onClick={() => navigate(`/expense/${expense.expense_id}`)}>Details</Button>
          }
        })
        return expensesListItem;
    }, [project, expenses])

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
        label: false
      };

      const pieShartTaxConfig = {
        appendPadding: 40,
        data: pieGraphTaxData,
        angleField: 'tax',
        colorField: 'expense_category',
        label: false
      };


      const expensesContentList: Record<string, React.ReactNode> = {
        expenses:  <Table size="small" style={{marginTop: '2px'}} columns={economicsColumns} dataSource={economicsData} />,
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
                        <div style={{width: '59%', marginRight: '1%', marginTop: '48px', boxShadow: cardShadow}}>
                        <Title style={{textAlign: 'center', marginTop: '24px'}} level={2}>Costs</Title>
                        <Pie {...pieShartConfig} />
                        </div>
                        <div style={{width: '59%', marginLeft: '1%', marginTop: '48px', boxShadow: cardShadow}}>
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
            bordered={false}
            style={{ borderRadius: 0, height: 'fit-content', boxShadow: 'none', padding: 0}}
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