/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useNavigate, useParams } from 'react-router-dom'
import { Button, Card, Space, Input, Typography, notification, Popconfirm, Divider, Select, Col, Row, Table } from 'antd';
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { updateProject } from '../../api/projects/update';
import { State } from '../../types/state';
import { deleteProject } from '../../api/projects/delete';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined } from "@ant-design/icons";
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { ProjectStatus, ProjectStatusTypes } from '../tags/ProjectStatus';
import { createProjectNote } from '../../api/notes/createProjectNote';
import { getAllProjectNotesByProjectId } from '../../api/notes/getAllProjectNotesByProductId';
import { deleteProjectNoteById } from '../../api/notes/deleteProjectNoteById';
import { cardShadow } from '../../enums/styles';
import { getAllProjectExpensesByProjectId } from '../../api/economics/expenses/getAllProjectExpensesByProjectId';
import { ExpenseObject } from '../../types/expense';
import { IncomeObject } from '../../types/income';
import * as React from 'react';
import { getAllProjectIncomesByProjectId } from '../../api/economics/incomes/getAllProjectIncomesByProjectId';

const { Text, Title, Link } = Typography;
const { TextArea } = Input;

const tabList = [
  {
    key: 'projectInformation',
    label: 'General information',
  },
  {
    key: 'projectEconomics',
    label: 'Economics',
  },
  {
    key: 'projectFiles',
    label: 'Files',
  },
];


const economicsTabList = [
  {
    key: 'expenses',
    label: 'Expenses',
  },
  {
    key: 'incomes',
    label: 'Incomes',
  },
];

const economicsColumns = [
  {
      title: 'Vendor',
      dataIndex: 'vendor',
      key: 'vendor'
  },
  {
      title: 'Description',
      dataIndex: 'description',
      key: 'description'
  },
  {
      title: 'Cost',
      dataIndex: 'cost',
      key: 'cost'
  },
  {
    title: 'Tax',
    dataIndex: 'tax',
    key: 'tax'
  },
  {
    title: '',
    dataIndex: 'operations',
    key: 'operations'
  },
];

const Project: React.FC = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useSelector((state : State) => state.user);
    const externalCompanies = useSelector((state : State) => state.application.externalCompanies);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const projects = useSelector((state : State) => state.application.projects);
    const [name, setName] = useState('');
    const [projectStatus, setProjectStatus] = useState<ProjectStatusTypes | ''>('');
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [editing, setEditing] = useState(false);
    const [projectNotes, setProjectNotes] = useState([]);
    const [projectExpenses, setProjectExpenses] = useState([]);
    const [projectIncomes, setProjectIncomes] = useState([]);
    const [activeTab, setActiveTab] = useState<string>('projectInformation');
    const [activeEconomicTab, setActiveEconomicTab] = useState<string>('expenses');
    const { id } =  useParams();
    const projectId = id || '';

    const project = projects.find((p : any) => p.id === projectId)

    useEffect(() => {
        if (project) {
          try {
            setName(project.name);
            setProjectStatus(project.status)
            } catch (error : any) {
              api.error({
                message: `Error fetching privilege`,
                description: error.toString(),
                placement: "bottom",
                duration: 1.4,
              });
            }
        }
        if (projectNotes && projectNotes.length === 0 && loggedInUser?.id) {
          getAllProjectNotesByProjectId(loggedInUser.id, projectId).then(response => {
            setProjectNotes(response)
          }).catch(error => {
            console.log('error fetching project notes', error)
          })
        }
        if (projectExpenses && projectExpenses.length === 0 && loggedInUser?.id) {
          getAllProjectExpensesByProjectId(loggedInUser.id, projectId).then(response => {
            if (!response.error && response.data) {
              setProjectExpenses(response.data)
            }
          }).catch(error => {
            console.log('error fetching project notes', error)
          })
        }
        if (projectIncomes && projectIncomes.length === 0 && loggedInUser?.id) {
          getAllProjectIncomesByProjectId(loggedInUser.id, projectId).then(response => {
            if (!response.error && response.data) {
              setProjectIncomes(response.data)
            }
          }).catch(error => {
            console.log('error fetching project notes', error)
          })
        }
      }, [projects]);

    const onHandleNameChange = (event : any) => setName(event.target.value);
    const onHandleProjectStatusChange = (value : any) => setProjectStatus(value);
    const onHandleNoteTitleChange = (event : any) => setNoteTitle(event.target.value);
    const onHandleNoteChange = (event : any) => setNote(event.target.value);
    const onHandleChangeActiveTab = (tab : string) => setActiveTab(tab);
    const onHandleChangeActiveEconomicTab = (tab: string) => setActiveEconomicTab(tab)

    const onSaveEdittedProject = () => {
        updateProject(loggedInUser.id, projectId, name, projectStatus)
        .then(response => {
          if (response?.error) {
            api.error({
                message: `Updated project failed`,
                description: response.message,
                placement: 'bottom',
                duration: 1.4
              });
              return
            }
            setEditing(false)
            api.info({
                message: `Updated project`,
                description: 'Succesfully updated project.',
                placement: 'bottom',
                duration: 1.4
            });
        })
        .catch(error => {
            api.error({
                message: `Error updating project`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    }

    const onClickdeleteProject = async () => {
        await deleteProject(loggedInUser.id, projectId)
          .then(response => {
            if (response?.error) {
              api.error({
                  message: `Deleted project failed`,
                  description: response.message,
                  placement: 'bottom',
                  duration: 1.4
                });
              return
            }
            api.info({
              message: `Deleted project`,
              description: "Succesfully deleted project.",
              placement: "bottom",
              duration: 1.2,
            });
            dispatch(popProject(projectId));
            setTimeout(() => {
              navigate("/projects");
            }, 1000);
          })
          .catch((error) => {
            api.error({
              message: `Error deleting project`,
              description: error.toString(),
              placement: "bottom",
              duration: 1.4,
            });
          });
      };

      const clearNoteFields = () => {
        setNoteTitle('');
        setNote('');
      }

      const onSubmitProjectNote = () => {
        const user = {
          id: loggedInUser.id,
          name: `${loggedInUser.firstName} ${loggedInUser.lastName}`,
          email: loggedInUser.email

        }
        createProjectNote(user, projectId, noteTitle, note).then(() => {
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

      const onClickdeleteNote = async (noteId : string, authorId : string, projectId : string) => {
        await deleteProjectNoteById(loggedInUser.id, noteId, authorId, projectId)
          .then(response => {
            if (response?.error) {
              api.error({
                  message: `Deleted project project note`,
                  description: response.message,
                  placement: 'bottom',
                  duration: 1.4
                });
              return
            }
            api.info({
              message: `Deleted project`,
              description: "Succesfully deleted project note.",
              placement: "bottom",
              duration: 1.2,
            });
          })
          .catch((error) => {
            api.error({
              message: `Error deleting project note`,
              description: error.toString(),
              placement: "bottom",
              duration: 1.4,
            });
          });
      };

      const getVendorName = (id : string) => externalCompanies.find(company => company.id === id)?.name;

      const expensesData: Array<any> = projectExpenses.map((expense : ExpenseObject) => {
        return {                    
            vendor: <Link href={`/income/${expense.vendor}`}>{getVendorName(expense.vendor)}</Link>,
            description: <Text>{expense.description}</Text>,
            cost: <Text>{expense.amount} {expense.currency}</Text>,
            tax: <Text>{expense.tax} {expense.currency}</Text>,
            operations: <Link href={`/expense/${expense.expense_id}`}>Details</Link>
           
          }
      })

      const incomesData: Array<any> = projectIncomes.map((income : IncomeObject) => {
        return {                    
            vendor:  <Link href={`/income/${income.vendor}`}>{getVendorName(income.vendor)}</Link>,
            description: <Text>{income.description}</Text>,
            cost: <Text>{income.amount} {income.currency}</Text>,
            tax: <Text>{income.tax} {income.currency}</Text>,
            operations: <Link href={`/income/${income.income_id}`}>Details</Link>
           
          }
      })

      const economicContentList: Record<string, React.ReactNode> = {
        expenses:  <Table size="small" style={{marginTop: '2px'}} columns={economicsColumns} dataSource={expensesData} />,
        incomes: <Table size="small" style={{marginTop: '2px'}} columns={economicsColumns} dataSource={incomesData} />,
      }

      const contentList: Record<string, React.ReactNode> = {
        projectInformation: (
          <div style={{padding: '24px'}}>
          <div style={{display: 'flex', justifyContent: 'flex-start', gap: '20px'}}>
          <Space direction="vertical" style={{minWidth: '320px'}}>
              <Text strong>Project name</Text>
              {editing ? (
                  <Input value={name} onChange={onHandleNameChange}/>
              ) : (
                  <Text>{name}</Text>
              )}
              <Text strong>Status</Text>
              {editing ? (
                <Select
                  style={{width: '100%'}}
                  options={[
                      {value: 'ongoing', label: 'ongoing'},
                      {value: 'cancelled', label: 'cancelled'},
                      {value: 'completed', label: 'completed'}
                  ]}
                  onChange={onHandleProjectStatusChange}
                  value={projectStatus}
                />
              ) : (
                  <ProjectStatus status={projectStatus ? projectStatus : 'default'} />
              )}
          </Space>
          <Space direction="vertical" style={{paddingRight: '24px'}}>
          <Text strong>Project ID</Text>
          <Text>{project?.id ? project.id : ''}</Text>
          <Text strong>Project started</Text>
          <Text>{project?.created_at ? project.created_at : ''}</Text>
          <Text strong>Project updated at</Text>
          <Text>{project?.updated_at ? project.updated_at : ''}</Text>
          </Space>
          </div>
          <Divider />
          <div style={{display: 'flex', justifyContent: 'flex-end', gap: '8px'}}>
              {editing && hasPrivilege(userPrivileges, PRIVILEGES.project_sudo) && (
                  <Popconfirm
                      placement="top"
                      title="Are you sure?"
                      description={`Do you want to delete user ${name}`}
                      onConfirm={onClickdeleteProject}
                      icon={<QuestionCircleOutlined style={{ color: "red" }} />}
                      okText="Yes"
                      cancelText="No"
                  >
                      <Button danger type="link">Delete</Button>
                  </Popconfirm>
              )}
              <Button type={editing ? "default" : "primary"} onClick={() => setEditing(!editing)}>{editing ? 'Close' : 'Edit'}</Button>
              {editing && (<Button type="primary" onClick={onSaveEdittedProject}>Save</Button>)}
          </div>
          </div>
        ),
        projectEconomics: (
          <Card 
            bordered={false}
            style={{ borderRadius: 0, marginTop: '2px'}}
            tabList={economicsTabList}
            activeTabKey={activeEconomicTab}
            bodyStyle={{padding: '0px', border: 'none'}}
            onTabChange={onHandleChangeActiveEconomicTab}
          >
          {economicContentList[activeEconomicTab]}
        </Card>),
        projectFiles: <p>Project files</p>,
      };

    return (
      <Row>
         {contextHolder}
      <Col span={16}>
        <Card 
          bordered={false}
          style={{ width: '98%', borderRadius: 0, height: 'fit-content', boxShadow: cardShadow, padding: 0}}
          tabList={tabList}
          activeTabKey={activeTab}
          bodyStyle={{padding: '0px'}}
          onTabChange={onHandleChangeActiveTab}
          >
          {contentList[activeTab]}
        </Card>
        </Col>
        <Col span={8}>
        <Card bordered={false} title="Create note" style={{ width: '100%', borderRadius: 0, height: 'fit-content', boxShadow: cardShadow}}>
          <Text strong style={{lineHeight: 2.3}}>Title</Text>
          <Input value={noteTitle} onChange={onHandleNoteTitleChange} />
          <Text strong style={{lineHeight: 2.3}}>Note</Text>
          <TextArea value={note} onChange={onHandleNoteChange}/>
          <div style={{display: 'flex', justifyContent: 'flex-end', gap: '8px', marginTop: '12px'}}>
            <Button onClick={clearNoteFields} type="link">Clear</Button>
            <Button type="primary" disabled={!note || !noteTitle} onClick={onSubmitProjectNote}>Submit</Button>
          </div>
        {projectNotes && (
        <div style={{paddingTop: '24px'}}>
        <Title level={5}>Notes</Title>
        <Divider style={{marginTop: '0px', marginBottom: '8px'}}/>
        {projectNotes.length > 0 && projectNotes.map((note : any) => (
          <div style={{width: '100%'}}>
            <div style={{display: 'flex', justifyContent: 'space-between', alignItems: 'center'}}>
              <Title level={5} style={{margin: '0px'}}>{note.title}</Title>
              <Popconfirm
                  placement="top"
                  title="Are you sure?"
                  description={`Do you want to delete note ${note.title}`}
                  onConfirm={() => onClickdeleteNote(note.id, note.author_id, note.project)}
                  icon={<QuestionCircleOutlined style={{ color: "red" }} />}
                  okText="Yes"
                  cancelText="No"
              >
                  <Button danger type="link">Delete note</Button>
              </Popconfirm>
            </div>
            <Text>{note.note}</Text>
            <div style={{display: 'flex', justifyContent: 'flex-end', flexDirection: 'column', marginTop: '4px'}}>
            <Text style={{textAlign: 'right', lineHeight: 1.2}}>{note.author_name}</Text>
            <Text style={{textAlign: 'right', lineHeight: 1.2}}>{note.author_email}</Text>
            <Text type="secondary" style={{textAlign: 'right'}}>{note.created_at}</Text>
            </div>
            <Divider style={{marginTop: '8px', marginBottom: '8px'}}/>
          </div>
          ))}
        </div>
        )}
        </Card>
        </Col>
      </Row>
      )
}

export default Project;