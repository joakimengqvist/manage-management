/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useParams } from 'react-router-dom'
import { Card, Typography, Row, Col, notification, Button, Table, Popconfirm, Modal, Select } from 'antd';
import { useEffect, useMemo, useState } from 'react';
import { getExternalCompanyById } from '../../api/externalCompanies/getById';
import { ExternalCompany } from '../../interfaces/externalCompany';
import { ExternalCompanyNote } from '../../interfaces/notes';
import CreateNote from '../notes/CreateNote';
import Notes from '../notes/Notes';
import { NOTE_TYPE } from '../../enums/notes';
import { createExternalCompanyNote } from '../../api/notes/externalCompany/create';
import { getAllExternalCompanyNotesByExternalCompanyId } from '../../api/notes/externalCompany/getAllByExternalCompanyId';
import { formatDateTimeToYYYYMMDDHHMM } from '../../helpers/stringDateFormatting';
import UpdateProjectExpense from './UpdateExternalCompany';
import { ExternalCompanyStatus } from '../status/ExternalCompanyStatus';
import { QuestionCircleOutlined, DeleteOutlined } from '@ant-design/icons';
import { useGetInvoices, useGetLoggedInUser, useGetUsers } from '../../hooks';
import { addInvoiceToCompany, getInvoicesByIds, removeInvoiceFromCompany } from '../../api';
import { Invoice } from '../../interfaces';
import { formatNumberWithSpaces } from '../../helpers/stringFormatting';
import InvoiceStatus from '../status/InvoiceStatus';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums';

const { Text, Title, Link } = Typography;

const externalCompanyTabList = [
    {
      key: 'info',
      label: 'Company information',
    },
    {
      key: 'invoices',
      label: 'Invoices',
    },
  ];

const invoiceColumns = [
    {
        title: 'Name',
        dataIndex: 'name',
        key: 'name'
    },
    {
        title: 'Price',
        dataIndex: 'price',
        key: 'price'
    },
    {
        title: 'Tax',
        dataIndex: 'tax',
        key: 'tax'
    },
    {
        title: 'Status',
        dataIndex: 'status',
        key: 'status'
    },
    {
        title: 'Due date',
        dataIndex: 'due_date',
        key: 'due_date'
    },
    {
        title: '',
        dataIndex: 'operations',
        key: 'operations'
    },
  ];


const ExternalCompanyDetails = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useGetLoggedInUser();
    const users = useGetUsers();
    const [externalCompanyNotes, setExternalCompanyNotes] = useState<Array<ExternalCompanyNote> | null>(null);
    const [invoices, setInvoices] = useState<Array<Invoice>>([]);
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [editing, setEditing] = useState(false);
    const [activeTab, setActiveTab] = useState('info');
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isModalLoading, setIsModalLoading] = useState(false);
    const [addedInvoiceToCompany, setAddedInvoiceToCompany] = useState('');
    const [invoiceOptions, setInvoiceOptions] = useState<Array<any>>([])
    const allInvoices = useGetInvoices();
    const [externalCompany, setExternalCompany] = useState<null | ExternalCompany>(null);
    const { id } =  useParams(); 
    const externalCompanyId = id || '';

    const getUserName = (id : string) => `${users?.[id]?.first_name} ${users?.[id]?.last_name}`;

    useEffect(() => {
        if (loggedInUser.id) {
            getExternalCompanyById(loggedInUser.id, externalCompanyId).then(response => {
                setExternalCompany(response.data)

                const invoiceSelectOptionsArray : Array<any> = []
                Object.keys(allInvoices).forEach(invoiceId => {
                    if (!response.data.invoices.includes(invoiceId)) {
                        invoiceSelectOptionsArray.push({
                            label: allInvoices[invoiceId].invoice_display_name,
                            value: allInvoices[invoiceId].id,
                        })
                    }
                })

                setInvoiceOptions(invoiceSelectOptionsArray)

                getInvoicesByIds(loggedInUser.id, response.data.invoices).then((response) => {
                    setInvoices(response.data)
                }).catch(error => {
                    console.log('error fetching', error)
                })
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
      const onHandleChangeAddInvoiceToCompany = (value : any) => setAddedInvoiceToCompany(value);
      const onHandleChangeActiveTab = (tab : string) => setActiveTab(tab);
  
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

        const onClickRemoveInvoiceFromCompany = (invoiceId : string) => {
            removeInvoiceFromCompany(loggedInUser.id, externalCompanyId, invoiceId).then((response) => {
                api.info({
                    message: response.message,
                    placement: "bottom",
                    duration: 1.2,
                });
            }).catch(error => {
                api.error({
                    message: `Error removing invoice`,
                    description: error.toString(),
                    placement: "bottom",
                    duration: 1.4,
                });
            })
        }

        const onAddInvoiceToCompany = () => {
            setIsModalLoading(true);
            addInvoiceToCompany(loggedInUser.id, externalCompanyId, addedInvoiceToCompany).then((response) => {
                api.info({
                    message: response.message,
                    placement: "bottom",
                    duration: 1.2,
                });
                setIsModalLoading(false);
            }).catch(error => {
                api.error({
                    message: `Error adding invoice`,
                    description: error.toString(),
                    placement: "bottom",
                    duration: 1.4,
                });
                setIsModalLoading(false);
            })
        }

        const invoiceData: Array<any> = useMemo(() => {
            const invoiceListItem = invoices && invoices.map((invoice : Invoice) => {
            return {                    
                name: <Link href={`/invoice/${invoice.id}`}>{invoice.invoice_display_name}</Link>,
                price: <Text>{formatNumberWithSpaces(invoice.actual_price)} SEK</Text>,
                tax: <Text>{formatNumberWithSpaces(invoice.actual_tax)} SEK</Text>,
                status: <InvoiceStatus status={invoice.status}/>,
                due_date: <Text>{formatDateTimeToYYYYMMDDHHMM(invoice.due_date)}</Text>,
                operations: (<>
                    {hasPrivilege(loggedInUser.privileges, PRIVILEGES.project_sudo) &&
                        <Popconfirm
                            placement="top"
                            title="Are you sure?"
                            description={`Do you want to remove invoice ${invoice.invoice_display_name} from ${externalCompany?.company_name}?`}
                            onConfirm={() => onClickRemoveInvoiceFromCompany(invoice.id)}
                            icon={<QuestionCircleOutlined twoToneColor="red" />}
                            okText="Yes"
                            cancelText="No"
                        >
                            <Button style={{padding: '4px'}} danger type="link"><DeleteOutlined /></Button>
                        </Popconfirm>
                    }
                </>)
              }
            })
            return invoiceListItem;
        }, [invoices])

        const externalCompanyContentList: Record<string, React.ReactNode> = {
            info:  (<>
            {externalCompany && (
                <div>
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
                </div>
                )}
            </>),
            invoices: (
                <div>
                    <div style={{width: '100%', display: 'flex', justifyContent: 'flex-end', padding: 8,}}>
                        <Button onClick={() => setIsModalOpen(true)}>Add invoice</Button>
                    </div>
                    <Table size="small" columns={invoiceColumns} dataSource={invoiceData} />
                </div>
            )
        }
        

    return (
        <Row>
            {contextHolder}
            <Col span={16} style={{ paddingRight: 8}}>
                <Card 
                    tabList={externalCompanyTabList}
                    activeTabKey={activeTab}
                    onTabChange={onHandleChangeActiveTab}
                    style={{ padding: 0}}
                    bodyStyle={activeTab === 'invoices' ? { padding: 0 } : {} }
                >
                    {externalCompanyContentList[activeTab]}
                </Card>
            </Col>
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
            <Modal
            open={isModalOpen}
            confirmLoading={isModalLoading}
            onCancel={() => setIsModalOpen(false)}
            footer={null}
            > 
                <Row>
                    <Col span={24}>
                        <Card style={{width: '100%'}}>
                            <Text>What invoice do you want to add to {externalCompany?.company_name}</Text>
                            <Select
                                style={{width: '100%', marginTop: '8px'}}
                                options={invoiceOptions}
                                onChange={onHandleChangeAddInvoiceToCompany}
                            />
                            <div style={{display: 'flex', justifyContent: 'flex-end', marginTop: '16px'}}>
                                <Button onClick={onAddInvoiceToCompany}>Add invoice</Button>
                            </div>
                        </Card>
                    </Col>
                </Row>
                <Row>
                    <Col span={24}>
                        <div style={{display: 'flex', justifyContent: 'flex-end', marginTop: '24px'}}>
                        <Button onClick={() => setIsModalOpen(false)}>Close</Button>
                        </div>
                    </Col>
                </Row>
            </Modal>
        </Row>
    )
}

export default ExternalCompanyDetails;