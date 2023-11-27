/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useParams } from 'react-router-dom'
import { Card, Typography, Row, Col, notification, Button, Divider } from 'antd';
import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../../interfaces/state';
import { ExpenseNote } from '../../../interfaces/notes';
import { ExpenseObject } from '../../../interfaces/expense'
import { getExpenseById } from '../../../api/economics/expenses/getById';
import { getAllExpenseNotesByExpenseId } from '../../../api/notes/expense/getAllByExpenseId';
import { createExpenseNote } from '../../../api/notes/expense/create'
import CreateNote from '../../notes/CreateNote';
import Notes from '../../notes/Notes'
import { NOTE_TYPE } from '../../../enums/notes';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import UpdateProjectExpense from './updateProjectExpense';
import ExpenseStatus from '../../status/ExpenseStatus';

const { Text, Title, Link } = Typography;

const Expense = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useSelector((state : State) => state.user);
    const [expense, setExpense] = useState<null | ExpenseObject>(null);
    const [expenseNotes, setExpenseNotes] = useState<Array<ExpenseNote> | null>(null);
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [editing, setEditing] = useState(false);
    const users = useSelector((state : State) => state.application.users);
    const externalCompanies = useSelector((state : State) => state.application.externalCompanies);
    const projects = useSelector((state : State) => state.application.projects);
    const { id } =  useParams(); 
    const expenseId = id || '';

    useEffect(() => {
        if (loggedInUser?.id) {
            getExpenseById(loggedInUser.id, expenseId).then(response => {
                setExpense(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
            getAllExpenseNotesByExpenseId(loggedInUser.id, expenseId).then(response => {
                setExpenseNotes(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
    }, [loggedInUser, expenseId]);

    const getUserName = (id : string) => {
        const user = users.find(user => user.id === id);
        return `${user?.first_name} ${user?.last_name}`;
    };
    const getVendorName = (id : string) => externalCompanies.find(company => company.id === id)?.company_name;
    const getProjectName = (id : string) => projects.find(project => project.id === id)?.name;

    const onHandleNoteTitleChange = (event : any) => setNoteTitle(event.target.value);
    const onHandleNoteChange = (event : any) => setNote(event.target.value);

    const clearNoteFields = () => {
    setNoteTitle('');
    setNote('');
    }

    const onSubmitExpenseNote = () => {
    const user = {
        id: loggedInUser.id,
        name: `${loggedInUser.firstName} ${loggedInUser.lastName}`,
        email: loggedInUser.email

    }
    createExpenseNote(user, expenseId, noteTitle, note).then((response) => {
        api.info({
            message: response.message,
            placement: "bottom",
            duration: 1.2,
        });
        }).catch(error => {
            api.error({
                message: `Error creating note`,
                description: error.toString(),
                placement: "bottom",
                duration: 1.4,
            });
        })
    }

    return (
        <Card>
            {expense && (
            <Row>
            {contextHolder}
            
            <Col span={15}>
                {!editing && (
                <Row>
                    <Col span={24}>
                        <div style={{display: 'flex', justifyContent: 'space-between'}}>
                            <span>
                                <Link style={{fontSize: '16px'}}href={`/external-company/${expense.vendor}`}>{getVendorName(expense.vendor)}</Link><br />
                                {expense.description}<br />
                                {formatDateTimeToYYYYMMDDHHMM(expense.expense_date)}<br />
                            </span>
                            <Button type="primary" onClick={() => setEditing(true)}>Edit expense info</Button>
                        </div>
                    </Col>
                    <Divider style={{marginTop: '16px', marginBottom: '16px'}}/>
                </Row>
                )}
                {editing ? (
                    <UpdateProjectExpense expense={expense} setEditing={setEditing} />
                ) : (
                <Row>
                    <Col span={8}  style={{padding: '0px 12px 12px 0px'}}>
                        <Text strong>Category</Text><br/>
                        {expense.expense_category}<br/>
                        <Text strong>Amount</Text><br/>
                        {`${expense.amount} ${expense.currency.toUpperCase()}`}<br/>
                        <Text strong>Tax</Text><br/>
                        {`${expense.tax} ${expense.currency.toUpperCase()}`}<br/>
                        <Text strong>Payment method</Text><br/>
                        {expense.payment_method}<br/>
                        <Text strong>Expense status</Text><br/>
                        <div style={{marginTop: '4px'}}>
                            <ExpenseStatus status={expense.status}/>
                        </div>
                    </Col>
                    <Col span={16} style={{padding: '0px 12px 12px 0px'}}>
                            <Text strong>Project</Text><br/>
                            <Link href={`/project/${expense.project_id}`}>{getProjectName(expense.project_id)}</Link><br/>
                            <Text strong>Created by</Text><br/>
                            <Link href={`/user/${expense.created_by}`}>{getUserName(expense.created_by)}</Link><br/>
                            <Text strong>Created at</Text><br/>
                            {formatDateTimeToYYYYMMDDHHMM(expense.created_at)}<br/>
                            <Text strong>Modified by</Text><br/>
                            <Link href={`/user/${expense.updated_by}`}>{getUserName(expense.updated_by)}</Link><br/>
                            <Text strong>Modified at</Text><br/>
                            {formatDateTimeToYYYYMMDDHHMM(expense.updated_at)}<br/>
                            
                    </Col>
                    <Divider style={{marginTop: '8px'}}/>
                </Row>
                )}
            </Col>
            <Col span={1}></Col>
            <Col span={8}>
                <Card>
                    <CreateNote
                        type={NOTE_TYPE.expense}
                        title={noteTitle}
                        onTitleChange={onHandleNoteTitleChange}
                        note={note}
                        onNoteChange={onHandleNoteChange}
                        onClearNoteFields={clearNoteFields}
                        onSubmit={onSubmitExpenseNote}
                    />
                    <Title level={4}>Notes</Title>
                    {expenseNotes && expenseNotes.length > 0 && 
                        <Notes notes={expenseNotes} type={NOTE_TYPE.expense} userId={loggedInUser.id} />
                    }
                </Card>
            </Col>
        </Row>
        )}
    </Card>
    )

}

export default Expense;