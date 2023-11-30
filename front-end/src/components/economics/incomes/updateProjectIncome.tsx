/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { Col, Divider, Row, Switch, Typography } from 'antd';
import { useDispatch } from 'react-redux';
import { Button, Input, Space, notification, DatePicker, Select } from 'antd';
import { appendPrivilege } from '../../../redux/applicationDataSlice';
import { IncomeAndExpenseCategoryOptions, IncomeAndExpenseCurrencyOptions, IncomeAndExpenseStatusOptions } from '../options';
import { IncomeObject } from '../../../interfaces/income';
import { updateIncome } from '../../../api/economics/incomes/update';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import { useGetExternalCompanies, useGetLoggedInUserId, useGetProjects } from '../../../hooks';

const { Text, Link } = Typography;
const { TextArea } = Input;

const numberPattern = /^[0-9]+$/;

const UpdateProjectIncome = ({ income, setEditing } : { income : IncomeObject, setEditing : (open : boolean) => void}) => {
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

    const getVendorName = (id : string) => externalCompanies?.[id]?.company_name;

    useEffect(() => {
        setProject(income.project_id); 
        setIncomeDate(income.income_date);
        setIncomeCategory(income.income_category);
        setIsStatisticsIncome(income.statistics_income);
        setVendor(income.vendor);
        setDescription(income.description);
        setAmount(income.amount.toString());
        setTax(income.tax.toString());
        setIncomeStatus(income.status.toString());
        setCurrency(income.currency);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

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
        updateIncome(
            income.id,
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

    const projectOptions = Object.keys(projects).map(projectId => ({ 
        label: projects[projectId].name, 
        value: projects[projectId].id
    }));

    const vendorOptions = Object.keys(externalCompanies).map(companyId => ({
        value: externalCompanies[companyId].id,
        label: externalCompanies[companyId].company_name
    }))

  return (<>
        {contextHolder}
        <Row>
            <Col span={24}>
                <div style={{display: 'flex', justifyContent: 'space-between'}}>
                    <span>
                        <Link style={{fontSize: '16px'}}href={`/external-company/${income.vendor}`}>{getVendorName(income.vendor)}</Link><br />
                        {income.description}<br />
                        {formatDateTimeToYYYYMMDDHHMM(income.income_date)}<br />
                    </span>
                    <div style={{display: 'flex', justifyContent: 'flex-end', gap: '12px', marginTop: '16px'}}>
                        <Button onClick={() => setEditing(false)}>Cancel</Button>
                        <Button type="primary" onClick={onSubmit}>Save</Button>
                    </div>
                </div>
            </Col>
            <Divider style={{marginTop: '16px', marginBottom: '16px'}}/>
        </Row>
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
                <div style={{display: 'flex', justifyContent: 'flex-end', alignItems: 'center', gap: '20px', marginBottom: '9px', marginTop: '30px'}}>
                        <Text strong>Is this a statistics income?</Text>
                        <Switch

                            checkedChildren={'Yes'}
                            unCheckedChildren={'No'}
                            checked={isStatisticsIncome}
                            onChange={() => setIsStatisticsIncome(isStatisticsIncome)}
                        />
                    </div>
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
    </>
  );
};

export default UpdateProjectIncome;