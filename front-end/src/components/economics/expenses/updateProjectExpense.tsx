/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { Col, Divider, Row, Typography } from 'antd';
import { useDispatch, useSelector } from 'react-redux';
import { Button, Input, Space, notification, DatePicker, Select } from 'antd';
import { State } from '../../../types/state';
import { appendPrivilege } from '../../../redux/applicationDataSlice';
import { IncomeAndExpenseCategoryOptions, IncomeAndExpenseCurrencyOptions, IncomeAndExpenseStatusOptions, paymentMethodOptions } from '../options';
import { ExpenseObject } from '../../../types/expense';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import { updateExpense } from '../../../api/economics/expenses/update';

const { Text, Link } = Typography;
const { TextArea } = Input;

const numberPattern = /^[0-9]+$/;

const UpdateProjectExpense = ({ expense, setEditing } : { expense : ExpenseObject, setEditing : (open : boolean) => void}) => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const allProjects = useSelector((state: State) => state.application.projects);
    const externalCompanies = useSelector((state: State) => state.application.externalCompanies);
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

    const getVendorName = (id : string) => externalCompanies.find(company => company.id === id)?.company_name;

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
    }, []);

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

    const vendorOptions = externalCompanies.map(company => ({
        value: company.id,
        label: company.company_name
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