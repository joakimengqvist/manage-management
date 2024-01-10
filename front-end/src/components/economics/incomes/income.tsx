/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useParams } from 'react-router-dom'
import { Card, Typography, Row, Col, notification, Button, Divider } from 'antd';
import { useEffect, useState } from 'react';
import { getIncomeById } from '../../../api/economics/incomes/getById';
import CreateNote from '../../notes/CreateNote';
import { createIncomeNote } from '../../../api/notes/income/create';
import { IncomeNote } from '../../../interfaces/notes';
import { Income } from '../../../interfaces/income'
import Notes from '../../notes/Notes';
import { NOTE_TYPE } from '../../../enums/notes';
import { getAllIncomeNotesByIncomeId } from '../../../api/notes/income/getAllByIncomeId';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import UpdateProjectIncome from './updateProjectIncome';
import IncomeStatus from '../../status/IncomeStatus';
import { GoldTag } from '../../tags/GoldTag';
import { useGetExternalCompanies, useGetLoggedInUser, useGetProjects, useGetUsers } from '../../../hooks';

const { Text, Title, Link } = Typography;

const IncomeDetails = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useGetLoggedInUser();
    const users = useGetUsers();
    const externalCompanies = useGetExternalCompanies();
    const projects = useGetProjects();
    const [income, setIncome] = useState<null | Income>(null);
    const [incomeNotes, setIncomeNotes] = useState<Array<IncomeNote> | null>(null);
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [editing, setEditing] = useState(false);
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
                setIncomeNotes(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUser, incomeId]);

      const getUserName = (id : string) => `${users?.[id]?.first_name} ${users?.[id]?.last_name}`;
      const getVendorName = (id : string) => externalCompanies?.[id]?.company_name || 'unknown';
      const getProjectName = (id : string) => projects?.[id]?.name;

      console.log('externalCompanies', externalCompanies);

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
        createIncomeNote(user, incomeId, noteTitle, note).then((response) => {
            api.info({
                message:response.message,
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

    return (<>
        {income && (
            <Row>
                {contextHolder}
                <Col span={16} style={{paddingRight: '8px'}}>
                    <div style={{display: 'flex', justifyContent: 'flex-end', gap: '12px', marginBottom: '6px'}}>
                    {editing ? (
                        <Button onClick={() => setEditing(false)}>Cancel</Button>
                    ) : (
                        <Button onClick={() => setEditing(true)}>Edit income info</Button>
                    )}    
                    </div>
                    <Card style={{marginBottom: '16px'}}>
                        {editing ? (
                            <UpdateProjectIncome income={income} />
                        ) : (
                        <>
                            <Row>
                                <Col span={24} style={{marginBottom: '0px'}}>
                                    <div style={{display: 'flex', justifyContent: 'space-between'}}>
                                        <div>
                                            <Link style={{fontSize: '16px'}}href={`/external-company/${income.vendor}`}>{getVendorName(income.vendor)}</Link><br />
                                            {income.description}<br />
                                            {formatDateTimeToYYYYMMDDHHMM(income.income_date)}<br />
                                        </div>
                                        <div style={{display: 'flex', flexDirection: 'column', alignItems: 'end', paddingRight: '0px', gap: '4px'}}>
                                            {income.statistics_income && <GoldTag label="Statistics income" />}
                                            <Link style={{paddingRight: '4px'}} href={`/invoice/${income.invoice_id}`}>Go to invoice</Link><br/>
                                        </div>
                                    </div>
                                </Col>
                                <Divider style={{marginBottom: '16px'}}/>
                            </Row>
                            <Row>
                                <Col span={8}  style={{padding: '0px 12px 12px 0px'}}>
                                    <Text strong>Category</Text><br/>
                                    {income.income_category}<br/>
                                    <Text strong>Amount</Text><br/>
                                    {`${income.amount} ${income.currency.toUpperCase()}`}<br/>
                                    <Text strong>Tax</Text><br/>
                                    {`${income.tax} ${income.currency.toUpperCase()}`}<br/>
                                    <Text strong>Income status</Text><br/>
                                    <div style={{marginTop: '4px'}}>
                                        <IncomeStatus status={income.status}/>
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
                                        <Link href={`/user/${income.updated_by}`}>{getUserName(income.updated_by)}</Link><br/>
                                        <Text strong>Modified at</Text><br/>
                                        {formatDateTimeToYYYYMMDDHHMM(income.updated_at)}<br/>
                                </Col>
                                <Divider style={{marginTop: '12px'}}/>
                            </Row>
                        </>)}
                        </Card>
                </Col>
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
    </>)
}

export default IncomeDetails;