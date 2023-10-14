/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useParams } from 'react-router-dom'
import { Card, Typography, Row, Col, notification, Button, Divider } from 'antd';
import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../../types/state';
import { getIncomeById } from '../../../api/economics/incomes/getIncomeById';
import CreateNote from '../../notes/CreateNote';
import { createIncomeNote } from '../../../api/notes/income/create';
import { IncomeNote } from '../../../types/notes';
import { IncomeObject } from '../../../types/income'
import Notes from '../../notes/Notes';
import { NOTE_TYPE } from '../../../enums/notes';
import { getAllIncomeNotesByIncomeId } from '../../../api/notes/income/getAllByIncomeId';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import { ExpenseAndIncomeStatus } from '../../tags/ExpenseAndIncomeStatus';

const { Text, Title, Link } = Typography;

const Income: React.FC = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useSelector((state : State) => state.user);
    const [income, setIncome] = useState<null | IncomeObject>(null);
    const [incomeNotes, setIncomeNotes] = useState<Array<IncomeNote> | null>(null);
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const users = useSelector((state : State) => state.application.users);
    const projects = useSelector((state : State) => state.application.projects);
    const externalCompanies = useSelector((state : State) => state.application.externalCompanies);
    const { id } =  useParams(); 
    const incomeId = id || '';

    useEffect(() => {
        if (loggedInUser?.id) {
            getIncomeById(loggedInUser.id, incomeId).then(response => {
                setIncome(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
            getAllIncomeNotesByIncomeId(loggedInUser.id, incomeId).then(response => {
                setIncomeNotes(response)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUser, incomeId]);

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

      const onSubmitIncomeNote = () => {
        const user = {
            id: loggedInUser?.id,
            name: `${loggedInUser?.firstName} ${loggedInUser?.lastName}`,
            email: loggedInUser?.email
    
        }
        createIncomeNote(user, incomeId, noteTitle, note).then(() => {
            api.info({
                message: `Created note`,
                description: "Succesfully created note.",
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
        <Card style={{ borderRadius: 0}}>
            {income && (
            <Row>
            {contextHolder}
            
            <Col span={15}>
                <Row>
                    <Col span={24} style={{marginBottom: '0px'}}>
                        <div style={{display: 'flex', justifyContent: 'space-between'}}>
                            <span>
                                <Link style={{fontSize: '16px'}}href={`/external-company/${income.vendor}`}>{getVendorName(income.vendor)}</Link><br />
                                {income.description}<br />
                                {formatDateTimeToYYYYMMDDHHMM(income.income_date)}<br />
                            </span>
                            <Button primary>Edit income info</Button>
                        </div>
                    </Col>
                    <Divider style={{marginTop: '16px', marginBottom: '16px'}}/>
                </Row>
                <Row>
                    <Col span={8}  style={{padding: '0px 12px 12px 0px'}}>
                        <Text strong>Category</Text><br/>
                        {income.income_category}<br/>
                        <Text strong>Amount</Text><br/>
                        {`${income.amount} ${income.currency.toUpperCase()}`}<br/>
                        <Text strong>Tax</Text><br/>
                        {`${income.tax} ${income.currency.toUpperCase()}`}<br/>
                        <Text strong>Payment method</Text><br/>
                        {income.payment_method}<br/>
                        <Text strong>Income status</Text><br/>
                        <div style={{marginTop: '4px'}}>
                            <ExpenseAndIncomeStatus status={income.status}/>
                        </div>
                    </Col>
                    <Col span={16} style={{padding: '0px 12px 12px 0px'}}>
                            <Text strong>Project</Text><br/>
                            <Link href={`/project/${income.project_id}`}>{getProjectName(income.project_id)}</Link><br/>
                            <Text strong>Created by</Text><br/>
                            <Link href={`/user/${income.created_by}`}>{getUserName(income.created_by)}</Link><br/>
                            <Text strong>Created at</Text><br/>
                            {formatDateTimeToYYYYMMDDHHMM(income.created_at)}<br/>
                            <Text strong>Modified by</Text><br/>
                            <Link href={`/user/${income.modified_by}`}>{getUserName(income.modified_by)}</Link><br/>
                            <Text strong>Modified at</Text><br/>
                            {formatDateTimeToYYYYMMDDHHMM(income.modified_at)}<br/>
                            <Text strong>Income ID</Text><br/>
                            {income.income_id}<br/>
                    </Col>
                    <Divider style={{marginTop: '12px'}}/>
                </Row>
            </Col>
            <Col span={1}></Col>
            <Col span={8}>
                <Card>
                    <CreateNote
                        type={NOTE_TYPE.income}
                        title={noteTitle}
                        onTitleChange={onHandleNoteTitleChange}
                        note={note}
                        onNoteChange={onHandleNoteChange}
                        onClearNoteFields={clearNoteFields}
                        onSubmit={onSubmitIncomeNote}
                    />
                      <Title level={4}>Notes</Title>
                        {incomeNotes && incomeNotes.length > 0 && 
                            <Notes notes={incomeNotes} type={NOTE_TYPE.income} userId={loggedInUser?.id} />
                        }
                        </Card>
            </Col>
        </Row>
        )}
    </Card>
    )

}

export default Income;