/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { Col, Divider, Row, Typography } from 'antd';
import { useDispatch } from 'react-redux';
import { Button, Input, Space, notification, DatePicker, Select } from 'antd';
import { appendPrivilege } from '../../../redux/applicationDataSlice';
import { IncomeAndExpenseCategoryOptions, IncomeAndExpenseCurrencyOptions, IncomeAndExpenseStatusOptions, paymentMethodOptions } from '../options';
import { Expense } from '../../../interfaces/expense';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import { updateExpense } from '../../../api/economics/expenses/update';
import { useGetExternalCompanies, useGetLoggedInUserId, useGetProjects } from '../../../hooks';

const { Text, Link } = Typography;
const { TextArea } = Input;

const numberPattern = /^[0-9]+$/;

const UpdateProjectExpense = ({ expense, setEditing } : { expense : Expense, setEditing : (open : boolean) => void}) => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUserId = useGetLoggedInUserId();
    const externalCompanies = useGetExternalCompanies();
    const projects = useGetProjects();
    const [project, setProject] = useState('');
    const [expenseDate, setExpenseDate] = useState('');
    const [expenseCategory, setExpenseCategory] = useState('');
    const [vendor, setVendor] = useState('');
    const [description, setDescription] = useState('');
    const [amount, setAmount] = useState('');
    const [tax, setTax] = useState('');
    const [expenseStatus, setExpenseStatus] = useState('');
    const [currency, setCurrency] = useState('');
    const [paymentMethod, setPaymentMethod] = useState('');

    const getVendorName = (id : string) => externalCompanies?.[id]?.company_name;

    useEffect(() => {
        setProject(expense.project_id); 
        setExpenseDate(expense.expense_date);
        setExpenseCategory(expense.expense_category);
        setVendor(expense.vendor);
        setDescription(expense.description);
        setAmount(expense.amount.toString());
        setTax(expense.tax.toString());
        setExpenseStatus(expense.status.toString());
        setCurrency(expense.currency);
        setPaymentMethod(expense.payment_method);
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

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

    const onChangeCurrency = (value : string) => setCurrency(value);
    const onChangePaymentMethod = (value : string) => setPaymentMethod(value);
    const onChangeExpenseCategory = (value : string) => setExpenseCategory(value); 
    const onChangeVendor = (value : string) => setVendor(value); 
    const onChangeProject = (value: string) => setProject(value);
    const onChangeExpenseStatus = (value : string) => setExpenseStatus(value);

    const onSubmit = () => {
        updateExpense(
            expense.id,
            project,
            expenseDate,
            expenseCategory,
            vendor,
            description,
            amount,
            tax,
            expenseStatus,
            currency,
            paymentMethod,
            loggedInUserId,
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
                message: `Error creating project expense`,
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
                        <Link style={{fontSize: '16px'}}href={`/external-company/${expense.vendor}`}>{getVendorName(expense.vendor)}</Link><br />
                        {expense.description}<br />
                        {formatDateTimeToYYYYMMDDHHMM(expense.expense_date)}<br />
                    </span>
                    <div style={{display: 'flex', justifyContent: 'flex-end', gap: '12px'}}>
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
                    <Text strong>Expense date</Text>
                    <DatePicker 
                        onChange={onChangeExpenseDate} 
                    />
                    <Text strong>Expense category</Text>
                    <Select
                        style={{width: '100%'}}
                        options={IncomeAndExpenseCategoryOptions}
                        onChange={onChangeExpenseCategory}
                        value={expenseCategory}
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
                    <Text strong>Expense status</Text>
                    <Select
                        placeholder="Select expense status"
                        style={{width: '100%'}}
                        options={IncomeAndExpenseStatusOptions}
                        onChange={onChangeExpenseStatus}
                        value={expenseStatus}
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
    </>
  );
};

export default UpdateProjectExpense;