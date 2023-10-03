/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { Col, Row, Typography } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { Button, Input, Space, Card, notification, DatePicker, Select } from 'antd';
import { State } from '../../../types/state';
import { appendPrivilege } from '../../../redux/applicationDataSlice';
import { cardShadow } from '../../../enums/styles';
import { createProjectExpense } from '../../../api/economics/expenses/createProjectExpense';

const { Title, Text } = Typography;
const { TextArea } = Input;

const numberPattern = /^[0-9]+$/;

const CreateProjectExpense: React.FC = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const allProjects = useSelector((state: State) => state.application.projects);
    const [project, setProject] = useState('');
    const [expenseDate, setExpenseDate] = useState('');
    const [expenseCategory, setExpenseCategory] = useState('');
    const [vendor, setVendor] = useState('');
    const [description, setDescription] = useState('');
    const [amount, setAmount] = useState('');
    const [tax, setTax] = useState('');
    const [currency, setCurrency] = useState('');
    const [paymentMethod, setPaymentMethod] = useState('');


    
    const projectOptions = allProjects.map(project => {
        return { label: project.name, value: project.id}
      }
    );

    const onChangeExpenseDate = (value : any) => {
        if (value) {
            setExpenseDate(value.$d)
        }
    }

    const onChangeAmount = (value : string) => {
        if (numberPattern.test(value)) {
            setAmount(value)
        }
    }

    const onChangeTaxAmount = (value : string) => {
        if (numberPattern.test(value)) {
            setTax(value)
        }
    }

    const onChangeCurrency = (value : any) => setCurrency(value);
    const onChangePaymentMethod = (value : any) => setPaymentMethod(value);
    const onChangeExpenseCategory = (value : any) => setExpenseCategory(value); 
    const onChangeProject = (value: any) => setProject(value);

    const onSubmit = () => {
        createProjectExpense(
            project,
            expenseDate,
            expenseCategory,
            vendor,
            description,
            amount,
            tax,
            currency,
            paymentMethod,
            userId,
        ).then(response => {
            if (response?.error || !response?.data) {
                api.error({
                    message: `Create project expense failed`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                    });
                return
            }
            api.info({
                message: `Created project expense`,
                description: 'Succesfully created project expense.',
                placement: 'bottom',
                duration: 1.4
            });
            dispatch(appendPrivilege({
                id: response.data,
                name: name,
                description: description
            }))
        })
        .catch(error => {
            api.error({
                message: `Error creating project expense`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    };

  return (
        <Card bordered={false} style={{borderRadius: 0, boxShadow: cardShadow, maxWidth: '600px'}}>
            {contextHolder}
            <Title level={4}>Create project expense</Title>
            <Row>
                <Col span={12} style={{padding: '12px 12px 12px 0px'}}>
                    <Space direction="vertical" style={{width: '100%'}}>
                        <Text strong>Project</Text>
                        <Select
                            placeholder="Project" 
                            style={{width: '100%'}}
                            options={projectOptions}
                            onChange={onChangeProject}
                            value={project}
                        />
                        <Text strong>Expense date</Text>
                        <DatePicker 
                            onChange={onChangeExpenseDate} 
                        />
                        <Text strong>Expense category</Text>
                        <Select
                            placeholder="Select payment method"
                            style={{width: '100%'}}
                            options={[
                                { value: 'materials', label: 'Materials' },
                                { value: 'fee', label: 'Fee' },
                                { value: 'staff', label: 'Staff' },
                                { value: 'equipment', label: 'Equipment' },
                                { value: 'rent', label: 'Rent' },
                                { value: 'travel', label: 'Travel' },
                                { value: 'marketing', label: 'Marketing' },
                                { value: 'maintenance', label: 'Maintenance' },
                                { value: 'software', label: 'Software' },
                                { value: 'consulting', label: 'Consulting' },
                                { value: 'insurance', label: 'Insurance' },
                                { value: 'utilities', label: 'Utilities' },
                                { value: 'advertising', label: 'Advertising' },
                            ]}
                            onChange={onChangeExpenseCategory}
                            value={expenseCategory}
                        />
                        <Text strong>Vendor</Text>
                        <Input 
                            placeholder="Vendor" 
                            value={vendor} 
                            onChange={event => setVendor(event.target.value)} 
                            onBlur={event => setVendor(event.target.value)}
                        />
                        <Text strong>Description</Text>
                        <TextArea 
                            placeholder="Description" 
                            value={description} 
                            onChange={(event : any) => setDescription(event.target.value)} 
                            onBlur={(event : any) => setDescription(event.target.value)}
                        />
                    </Space>
                </Col>
                <Col span={12} style={{padding: '12px'}}>
                    <Space direction="vertical" style={{width: '100%'}}>
                        <Text strong>Currency</Text>
                        <Select
                            placeholder="Select currency"
                            style={{width: '100%'}}
                            options={[
                                { value: 'sek', label: 'SEK' },
                                { value: 'nok', label: 'NOK' },
                                { value: 'us', label: 'US' },
                                { value: 'eur', label: 'EUR' },
                            ]}
                            onChange={onChangeCurrency}
                            value={currency}
                        />
                        <Text strong>Amount</Text>
                        <Input 
                            placeholder="Amount" 
                            value={amount} 
                            onChange={event => onChangeAmount(event.target.value)} 
                            onBlur={event => onChangeAmount(event.target.value)}
                            suffix={currency.toUpperCase()}
                        />
                        <Text strong>Tax</Text>
                        <Input 
                            placeholder="Tax" 
                            value={tax} 
                            onChange={event => onChangeTaxAmount(event.target.value)} 
                            onBlur={event => onChangeTaxAmount(event.target.value)}
                            suffix={currency.toUpperCase()}
                        />
                        <Text strong>Payment method</Text>
                        <Select
                            placeholder="Select payment method"
                            style={{width: '100%'}}
                            options={[
                                { value: 'debit-card', label: 'Debit card' },
                                { value: 'credit-card', label: 'Credit card' },
                                { value: 'invoice', label: 'Invoice' },
                                { value: 'paypal', label: 'PayPal' },
                                { value: 'bank-transfer', label: 'Bank Transfer' },
                                { value: 'check', label: 'Check' },
                                { value: 'cash', label: 'Cash' },
                                { value: 'crypto', label: 'Cryptocurrency' },
                                { value: 'mobile-wallet', label: 'Mobile Wallet' },
                                { value: 'apple-pay', label: 'Apple Pay' },
                                { value: 'google-pay', label: 'Google Pay' },
                                { value: 'amazon-pay', label: 'Amazon Pay' },
                            ]}
                            onChange={onChangePaymentMethod}
                            value={paymentMethod}
                        />
                    </Space>
                </Col>
            </Row>
            <Row>
                <Col>
                    <div style={{display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px'}}>
                        <Button type="primary" onClick={onSubmit}>Create expense</Button>
                    </div>
                </Col>
            </Row>
        </Card>
  );
};

export default CreateProjectExpense;