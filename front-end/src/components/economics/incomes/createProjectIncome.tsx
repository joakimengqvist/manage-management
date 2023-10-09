/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { Col, Row, Typography } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { Button, Input, Space, Card, notification, DatePicker, Select } from 'antd';
import { State } from '../../../types/state';
import { appendPrivilege } from '../../../redux/applicationDataSlice';
import { cardShadow } from '../../../enums/styles';
import { createProjectIncome } from '../../../api/economics/incomes/createProjectIncome';
import { IncomeAndExpenseCategoryOptions, IncomeAndExpenseCurrencyOptions, IncomeAndExpenseStatusOptions, paymentMethodOptions } from '../options';

const { Title, Text } = Typography;
const { TextArea } = Input;

const numberPattern = /^[0-9]+$/;

const CreateProjectIncome: React.FC = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const allProjects = useSelector((state: State) => state.application.projects);
    const externalCompanies = useSelector((state: State) => state.application.externalCompanies);
    const [project, setProject] = useState('');
    const [incomeDate, setIncomeDate] = useState('');
    const [incomeCategory, setIncomeCategory] = useState('');
    const [vendor, setVendor] = useState('');
    const [description, setDescription] = useState('');
    const [amount, setAmount] = useState('');
    const [tax, setTax] = useState('');
    const [incomeStatus, setIncomeStatus] = useState('');
    const [currency, setCurrency] = useState('');
    const [paymentMethod, setPaymentMethod] = useState('');


    
    const projectOptions = allProjects.map(project => {
        return { label: project.name, value: project.id}
      }
    );

    const onChangeIncomeDate = (value : any) => {
        if (value) {
            setIncomeDate(value.$d)
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

    const onChangeCurrency = (value : string) => setCurrency(value);
    const onChangePaymentMethod = (value : string) => setPaymentMethod(value);
    const onChangeIncomeCategory = (value : string) => setIncomeCategory(value); 
    const onChangeVendor = (value : string) => setVendor(value); 
    const onChangeProject = (value: string) => setProject(value);
    const onChangeIncomeStatus = (value : string) => setIncomeStatus(value);

    const onSubmit = () => {
        createProjectIncome(
            project,
            incomeDate,
            incomeCategory,
            vendor,
            description,
            amount,
            tax,
            incomeStatus,
            currency,
            paymentMethod,
            userId,
        ).then(response => {
            if (response?.error || !response?.data) {
                api.error({
                    message: `Create project income failed`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                    });
                return
            }
            api.info({
                message: `Created project income`,
                description: 'Succesfully created project income.',
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
                message: `Error creating project income`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    };

    const vendorOptions = externalCompanies.map(company => ({
        value: company.id,
        label: company.name
    }))

  return (
        <Card bordered={false} style={{borderRadius: 0, boxShadow: cardShadow, maxWidth: '600px'}}>
            {contextHolder}
            <Title level={4}>Create project income</Title>
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
                        <Text strong>Income date</Text>
                        <DatePicker 
                            onChange={onChangeIncomeDate} 
                        />
                        <Text strong>Income category</Text>
                        <Select
                            style={{width: '100%'}}
                            options={IncomeAndExpenseCategoryOptions}
                            onChange={onChangeIncomeCategory}
                            value={incomeCategory}
                        />
                        <Text strong>Vendor</Text>
                        <Select 
                            style={{width: '100%'}}
                            value={vendor} 
                            onChange={onChangeVendor} 
                            options={vendorOptions}
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
                            options={IncomeAndExpenseCurrencyOptions}
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
                        <Text strong>Income status</Text>
                        <Select
                            placeholder="Select income status"
                            style={{width: '100%'}}
                            options={IncomeAndExpenseStatusOptions}
                            onChange={onChangeIncomeStatus}
                            value={incomeStatus}
                        />
                        <Text strong>Payment method</Text>
                        <Select
                            placeholder="Select payment method"
                            style={{width: '100%'}}
                            options={paymentMethodOptions}
                            onChange={onChangePaymentMethod}
                            value={paymentMethod}
                        />
                    </Space>
                </Col>
            </Row>
            <Row>
                <Col>
                    <div style={{display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px'}}>
                        <Button type="primary" onClick={onSubmit}>Create income</Button>
                    </div>
                </Col>
            </Row>
        </Card>
  );
};

export default CreateProjectIncome;