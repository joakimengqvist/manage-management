/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useParams } from 'react-router-dom'
import { Card, Typography, Row, Col, notification, Button } from 'antd';
import { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { getExternalCompanyById } from '../../api/externalCompanies/getById';
import { ExternalCompany } from '../../types/externalCompany';
import { ExternalCompanyNote } from '../../types/notes';
import CreateNote from '../notes/CreateNote';
import Notes from '../notes/Notes';
import { NOTE_TYPE } from '../../enums/notes';
import { createExternalCompanyNote } from '../../api/notes/externalCompany/create';
import { getAllExternalCompanyNotesByExternalCompanyId } from '../../api/notes/externalCompany/getAllByExternalCompanyId';
import { formatDateTimeToYYYYMMDDHHMM } from '../../helpers/stringDateFormatting';
import UpdateProjectExpense from './UpdateExternalCompany';
import { ExternalCompanyStatus } from '../status/ExternalCompanyStatus';

const { Text, Title, Link } = Typography;

const ExternalCompanyDetails: React.FC = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useSelector((state : State) => state.user);
    const users = useSelector((state : State) => state.application.users);
    const [externalCompanyNotes, setExternalCompanyNotes] = useState<Array<ExternalCompanyNote> | null>(null);
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [editing, setEditing] = useState(false);
    const [externalCompany, setExternalCompany] = useState<null | ExternalCompany>(null);
    const { id } =  useParams(); 
    const externalCompanyId = id || '';

    const getUserName = (userId : string) => users.find(user => user.id === userId)?.first_name;

    useEffect(() => {
        if (loggedInUser.id) {
            getExternalCompanyById(loggedInUser.id, externalCompanyId).then(response => {
                setExternalCompany(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
            getAllExternalCompanyNotesByExternalCompanyId(loggedInUser.id, externalCompanyId).then(response => {
                setExternalCompanyNotes(response.data)
            }).catch(error => {
                console.log('error fetching', error)
            })
        }
      }, [loggedInUser.id]);

      const onHandleNoteTitleChange = (event : any) => setNoteTitle(event.target.value);
      const onHandleNoteChange = (event : any) => setNote(event.target.value);
  
      const clearNoteFields = () => {
        setNoteTitle('');
        setNote('');
      }

      const onSubmitIncomeNote = () => {
        const user = {
            id: loggedInUser.id,
            name: `${loggedInUser.firstName} ${loggedInUser.lastName}`,
            email: loggedInUser.email
    
        }
        createExternalCompanyNote(user, externalCompanyId, noteTitle, note).then((response) => {
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

    return (<>
        <Card style={{ padding: 0}}>
            {contextHolder}
                {externalCompany && (
                    <Row>
                        <Col span={15}>
                            {!editing && (
                            <Row>
                                <Col span={24}>
                                    <div style={{display: 'flex', justifyContent: 'space-between'}}>
                                        <div>
                                            <Title style={{marginBottom: '8px'}} level={4}>{externalCompany.company_name}</Title>
                                            <ExternalCompanyStatus status={externalCompany.status} />
                                        </div>
                                        <Button type="primary" onClick={() => setEditing(true)}>Edit company info</Button>
                                    </div>
                                </Col>
                            </Row>
                            )}
                            {editing ? (
                                <UpdateProjectExpense externalCompany={externalCompany} setEditing={setEditing} />
                            ) : (
                            <Row>
                                <Col span={7}  style={{padding: '0px 12px 12px 0px'}}>
                                    <Title style={{marginTop: '12px', marginBottom: '4px'}} level={5}>Address info</Title>
                                    {`${externalCompany.country}, ${externalCompany.city}`}<br />
                                    {externalCompany.address}<br />
                                    {`${externalCompany.postal_code} ${externalCompany.state_province}`}
                                    <Title style={{marginTop: '12px', marginBottom: '4px'}} level={5}>Contact info</Title>
                                    {externalCompany.contact_person}<br />
                                    {externalCompany.contact_email}<br />
                                    {externalCompany.contact_phone}
                                </Col>
                                <Col span={17} style={{padding: '0px 12px 12px 0px'}}>
                                    <Title style={{marginTop: '12px', marginBottom: '4px'}} level={5}>Company info</Title>
                                    <div style={{display: 'flex', marginBottom: '2px'}}>
                                        <Text strong style={{minWidth: '200px'}}>Registration number:</Text><Text>{externalCompany.company_registration_number}</Text>
                                    </div>
                                    <div style={{display: 'flex', marginBottom: '2px'}}>
                                        <Text strong style={{minWidth: '200px'}}>Tax identification number:</Text><Text>{externalCompany.tax_identification_number}</Text>
                                    </div>
                                    <div style={{display: 'flex', marginBottom: '2px'}}>
                                        <Text strong style={{minWidth: '200px'}}>Billing currency:</Text><Text>{externalCompany.billing_currency}</Text>
                                    </div>
                                    <div style={{display: 'flex', marginBottom: '2px'}}>
                                        <Text strong style={{minWidth: '200px'}}>Bank account info:</Text><Text>{externalCompany.bank_account_info}</Text>
                                    </div>
                                    <Title style={{marginTop: '12px', marginBottom: '4px'}} level={5}>Other info</Title>
                                    <div style={{display: 'flex', marginBottom: '2px'}}>
                                        <Text strong style={{minWidth: '200px'}}>Created at:</Text><Text>{formatDateTimeToYYYYMMDDHHMM(externalCompany.created_at)}</Text>
                                    </div>
                                    <div style={{display: 'flex', marginBottom: '2px'}}>
                                        <Text strong style={{minWidth: '200px'}}>Created by:</Text><Link href={`/user/${externalCompany.created_by}`}>{getUserName(externalCompany.created_by)}</Link>
                                    </div>
                                    <div style={{display: 'flex', marginBottom: '2px'}}>
                                        <Text strong style={{minWidth: '200px'}}>Updated at:</Text><Text>{formatDateTimeToYYYYMMDDHHMM(externalCompany.updated_at)}</Text>
                                    </div>
                                    <div style={{display: 'flex', marginBottom: '2px'}}>
                                        <Text strong style={{minWidth: '200px'}}>Updated by:</Text><Link href={`/user/${externalCompany.created_by}`}>{getUserName(externalCompany.updated_by)}</Link>
                                    </div>
                                    
                                </Col>
                            </Row>
                            )}
                        </Col>
                       
                        <Col span={1}></Col>
                        <Col span={8}>
                            <Card>
                                <CreateNote
                                    type={NOTE_TYPE.external_company}
                                    title={noteTitle}
                                    onTitleChange={onHandleNoteTitleChange}
                                    note={note}
                                    onNoteChange={onHandleNoteChange}
                                    onClearNoteFields={clearNoteFields}
                                    onSubmit={onSubmitIncomeNote}
                                />
                                {externalCompanyNotes && externalCompanyNotes.length > 0 && 
                                    <Notes notes={externalCompanyNotes} type={NOTE_TYPE.external_company} userId={loggedInUser.id} />
                                }
                            </Card>
                        </Col>
                    </Row> 
                )}
        </Card>
    </>)

}

export default ExternalCompanyDetails;