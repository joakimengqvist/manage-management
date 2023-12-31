/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { Checkbox, Col, Row, Typography } from 'antd';
import { useDispatch } from 'react-redux';
import { Button, Input, Space, Card, notification, DatePicker, Select } from 'antd';
import { appendPrivilege } from '../../../redux/applicationDataSlice';
import { createIncome } from '../../../api/economics/incomes/create';
import { IncomeAndExpenseCategoryOptions, IncomeAndExpenseCurrencyOptions, IncomeAndExpenseStatusOptions } from '../options';
import { useGetExternalCompanies, useGetLoggedInUserId, useGetProjects } from '../../../hooks';

const { Title, Text } = Typography;
const { TextArea } = Input;

const numberPattern = /^[0-9]+$/;

const CreateProjectIncome = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUserId = useGetLoggedInUserId();
    const externalCompanies = useGetExternalCompanies();
    const projects = useGetProjects();
    const [project, setProject] = useState('');
    const [incomeDate, setIncomeDate] = useState('');
    const [incomeCategory, setIncomeCategory] = useState('');
    const [isStatisticsIncome, setIsStatisticsIncome] = useState(false);
    const [vendor, setVendor] = useState('');
    const [description, setDescription] = useState('');
    const [amount, setAmount] = useState('');
    const [tax, setTax] = useState('');
    const [incomeStatus, setIncomeStatus] = useState('');
    const [currency, setCurrency] = useState('');

    const projectOptions = Object.keys(projects).map(projectId => ({
        label: projects?.[projectId]?.name, 
        value: projects?.[projectId]?.id
    }));

    const vendorOptions = Object.keys(externalCompanies).map(companyId => ({
        value: externalCompanies[companyId].id,
        label: externalCompanies[companyId].company_name
    }))

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
    const onChangeIncomeCategory = (value : string) => setIncomeCategory(value); 
    const onChangeVendor = (value : string) => setVendor(value); 
    const onChangeProject = (value: string) => setProject(value);
    const onChangeIncomeStatus = (value : string) => setIncomeStatus(value);

    const onSubmit = () => {
        createIncome(
            project,
            incomeDate,
            incomeCategory,
            isStatisticsIncome,
            vendor,
            description,
            amount,
            tax,
            incomeStatus,
            currency,
            loggedInUserId,
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
                message: response.message,
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

  return (
        <Card style={{maxWidth: '600px'}}>
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
                        <Text strong>Is this a statistical income?</Text>
                        <Checkbox
                            checked={isStatisticsIncome}
                            onChange={(event : any) => setIsStatisticsIncome(event.target.checked)}
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