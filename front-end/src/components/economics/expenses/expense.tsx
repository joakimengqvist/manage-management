/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useParams } from 'react-router-dom'
import { Card, Space, Typography, Row, Col } from 'antd';
import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../../types/state';
import { cardShadow } from '../../../enums/styles';
import { getExpenseById } from '../../../api/economics/expenses/getExpenseById';

const { Text, Title, Link } = Typography;

type ExpenseObject = {
	expense_id: string,
	project_id: string,
    expense_date: any,
	expense_category: string,
	vendor: string,
	description: string,
	amount: number,
	tax: number,
    status: string,
	currency: string,
	payment_method: string,
	created_by: string,
	created_at: any,
	modified_by: any,
	modified_at: any
}

const Expense: React.FC = () => {
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const [expense, setExpense] = useState<null | ExpenseObject>(null);
    const externalCompanies = useSelector((state : State) => state.application.externalCompanies);
    const { id } =  useParams(); 
    const expenseId = id || '';

    useEffect(() => {
        if (loggedInUserId) {
            getExpenseById(loggedInUserId, expenseId).then(response => {
                setExpense(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUserId]);

      const getVendorName = (id : string) => externalCompanies.find(company => company.id === id)?.name;

    return (
        <Card bordered={false} style={{ borderRadius: 0, boxShadow: cardShadow}}>
            <Title level={4}>Expense information</Title>
                {expense && (
                    <Row>
                        <Col span={8}>
                            <Space direction="vertical">
                                <Text strong>Vendor</Text>
                                <Link href={`/external-company/${expense.vendor}`}>{getVendorName(expense.vendor)}</Link>
                                <Text strong>Description</Text>
                                {expense.description}
                                <Text strong>Expense date</Text>
                                {expense.expense_date}
                                <Text strong>Expense category</Text>
                                {expense.expense_date}
                            </Space>
                        </Col>
                        <Col span={8}>
                            <Space direction="vertical">
                                <Text strong>Amount</Text>
                                {`${expense.amount} ${expense.currency.toUpperCase()}`}
                                <Text strong>Tax</Text>
                                {`${expense.tax} ${expense.currency.toUpperCase()}`}
                                <Text strong>Payment method</Text>
                                {expense.payment_method}
                                <Text strong>Expense status</Text>
                                {expense.status}
                            </Space>
                        </Col>
                        <Col span={8}>
                            <Space direction="vertical">
                            <Text strong>Expense ID</Text>
                                {expense.expense_id}
                                <Text strong>Project ID</Text>
                                {expense.project_id}
                                <Text strong>Created by</Text>
                                {expense.created_by}
                                <Text strong>Created at</Text>
                                {expense.created_at}
                                <Text strong>Modified by</Text>
                                {expense.modified_by}
                                <Text strong>Modified at</Text>
                                {expense.modified_at}
                            </Space>
                        </Col>
                    </Row>
                )}
        </Card>
    )

}

export default Expense;