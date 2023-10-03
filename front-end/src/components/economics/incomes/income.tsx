/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useParams } from 'react-router-dom'
import { Card, Space, Typography, Row, Col } from 'antd';
import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../../types/state';
import { cardShadow } from '../../../enums/styles';
import { getIncomeById } from '../../../api/economics/incomes/getIncomeById';

const { Text, Title } = Typography;

type IncomeObject = {
	income_id: string,
	project_id: string,
    income_date: any,
	income_category: string,
	vendor: string,
	description: string,
	amount: number,
	tax: number,
	currency: string,
	payment_method: string,
	created_by: string,
	created_at: any,
	modified_by: any,
	modified_at: any
}

const Income: React.FC = () => {
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const [income, setIncome] = useState<null | IncomeObject>(null);
    const { id } =  useParams(); 
    const incomeId = id || '';

    useEffect(() => {
        if (loggedInUserId) {
            getIncomeById(loggedInUserId, incomeId).then(response => {
                setIncome(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUserId]);

    return (
        <Card bordered={false} style={{ borderRadius: 0, boxShadow: cardShadow}}>
            <Title level={4}>Income information</Title>
                {income && (
                    <Row>
                        <Col span={8}>
                            <Space direction="vertical">
                                <Text strong>Vendor</Text>
                                {income.vendor}
                                <Text strong>Description</Text>
                                {income.description}
                                <Text strong>Income date</Text>
                                {income.income_date}
                                <Text strong>Income category</Text>
                                {income.income_date}
                            </Space>
                        </Col>
                        <Col span={8}>
                            <Space direction="vertical">
                                <Text strong>Amount</Text>
                                {`${income.amount} ${income.currency.toUpperCase()}`}
                                <Text strong>Tax</Text>
                                {`${income.tax} ${income.currency.toUpperCase()}`}
                                <Text strong>Payment method</Text>
                                {income.payment_method}

                            </Space>
                        </Col>
                        <Col span={8}>
                            <Space direction="vertical">
                            <Text strong>Income ID</Text>
                                {income.income_id}
                                <Text strong>Project ID</Text>
                                {income.project_id}
                                <Text strong>Created by</Text>
                                {income.created_by}
                                <Text strong>Created at</Text>
                                {income.created_at}
                                <Text strong>Modified by</Text>
                                {income.modified_by}
                                <Text strong>Modified at</Text>
                                {income.modified_at}
                            </Space>
                        </Col>
                    </Row>
                )}
        </Card>
    )

}

export default Income;