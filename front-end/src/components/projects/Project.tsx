/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useNavigate, useParams } from 'react-router-dom'
import { Button, Card, Input, Typography, notification, Popconfirm, Divider, Select, Col, Row, Table } from 'antd';
import { useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';
import { updateProject } from '../../api/projects/update';
import { deleteProject } from '../../api/projects/delete';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined, DeleteOutlined, ZoomInOutlined } from "@ant-design/icons";
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { createProjectNote } from '../../api/notes/project/create';
import { getAllProjectNotesByProjectId } from '../../api/notes/project/getAllByProjectId';
import { getAllExpensesByProjectId } from '../../api/economics/expenses/getAllByProjectId';
import { ExpenseObject } from '../../interfaces/expense';
import { IncomeObject } from '../../interfaces/income';
import * as React from 'react';
import { getAllIncomesByProjectId } from '../../api/economics/incomes/getAllByProjectId';
import CreateNote from '../notes/CreateNote';
import { formatDateTimeToYYYYMMDDHHMM } from '../../helpers/stringDateFormatting';
import { NOTE_TYPE } from '../../enums/notes';
import NoteList from '../notes/Notes';
import ProjectStatus from '../status/ProjectStatus';
import { Project, ProjectNote } from '../../interfaces';
import { getProjectById } from '../../api';
import { useGetExternalCompanies, useGetLoggedInUser, useGetProjects, useGetUsers } from '../../hooks';

const { Text, Link } = Typography;

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

const ProjectDetails = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useGetLoggedInUser();
    const users = useGetUsers();
    const projects = useGetProjects();
    const externalCompanies = useGetExternalCompanies();
    const [project, setProject]= useState<Project | null>(null);
    const [name, setName] = useState('');
    const [projectStatus, setProjectStatus] = useState('');
    const [noteTitle, setNoteTitle] = useState('');
    const [note, setNote] = useState('');
    const [editing, setEditing] = useState(false);
    const [projectNotes, setProjectNotes] = useState<Array<ProjectNote>>([]);
    const [projectExpenses, setProjectExpenses] = useState<Array<ExpenseObject>>([]);
    const [projectIncomes, setProjectIncomes] = useState<Array<IncomeObject>>([]);
    const [activeTab, setActiveTab] = useState<string>('projectInformation');
    const [activeEconomicTab, setActiveEconomicTab] = useState<string>('expenses');
    const { id } =  useParams();
    const projectId = id || '';

    useEffect(() => {
      getProjectById(loggedInUser.id, projectId).then(response => {
        if (response?.error) {
          api.error({
            message: `Error fetching project`,
            description: response.message,
            placement: "bottom",
            duration: 1.4,
          })
        }
        setName(response.data.name);
        setProjectStatus(response.data.status)
        setProject(response.data)
        }).catch(error => {
          console.log('error fetching project', error)
        })

        if (projectNotes && projectNotes.length === 0 && loggedInUser?.id) {
          getAllProjectNotesByProjectId(loggedInUser.id, projectId).then(response => {
            if (!response.error && response.data) {
              setProjectNotes(response.data)
            }
          }).catch(error => {
            console.log('error fetching project notes', error)
          })
        }

        if (projectExpenses && projectExpenses.length === 0 && loggedInUser?.id) {
          getAllExpensesByProjectId(loggedInUser.id, projectId).then(response => {
            if (!response.error && response.data) {
              setProjectExpenses(response.data)
            }
          }).catch(error => {
            console.log('error fetching project notes', error)
          })
        }

        if (projectIncomes && projectIncomes.length === 0 && loggedInUser?.id) {
          getAllIncomesByProjectId(loggedInUser.id, projectId).then(response => {
            if (!response.error && response.data) {
              setProjectIncomes(response.data)
            }
          }).catch(error => {
            console.log('error fetching project notes', error)
          })
        }

      }, [projects, projectId, projectNotes, projectExpenses, projectIncomes]);

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
                message: response.message,
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
              message: response.message,
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
        createProjectNote(user, projectId, noteTitle, note).then((response) => {
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

      const getUserName = (id : string) => `${users?.[id]?.first_name} ${users?.[id]?.last_name}`;
      const getVendorName = (id : string) => externalCompanies?.[id]?.company_name;

      const expensesData: Array<any> = projectExpenses.map((expense : ExpenseObject) => {
        return {                    
            vendor: <Link href={`/external-company/${expense?.vendor}`}>{getVendorName(expense?.vendor)}</Link>,
            description: <Text>{expense?.description}</Text>,
            cost: <Text>{expense?.amount} {expense?.currency}</Text>,
            tax: <Text>{expense?.tax} {expense?.currency}</Text>,
            operations: <Link href={`/expense/${expense?.id}`}><ZoomInOutlined /></Link>
           
          }
      })

      const incomesData: Array<any> = projectIncomes.map((income : IncomeObject) => {
        return {                    
            vendor:  <Link href={`/external-company/${income.vendor}`}>{getVendorName(income.vendor)}</Link>,
            description: <Text>{income.description}</Text>,
            cost: <Text>{income.amount} {income.currency}</Text>,
            tax: <Text>{income.tax} {income.currency}</Text>,
            operations: <Link href={`/income/${income.id}`}><ZoomInOutlined /></Link>
           
          }
      })

      const economicContentList: Record<string, React.ReactNode> = {
        expenses:  <Table size="small" columns={economicsColumns} dataSource={expensesData} />,
        incomes: <Table size="small" columns={economicsColumns} dataSource={incomesData} />,
      }

      const contentList: Record<string, React.ReactNode> = {
        projectInformation: (
          <div style={{padding: '24px'}}>
          <div style={{display: 'flex', justifyContent: 'flex-start', gap: '20px'}}>
          <div style={{minWidth: '320px'}}>
              <div style={{paddingBottom: '4px'}}>
                <Text strong>Project name</Text><br />
                {editing ? (
                    <Input value={name} onChange={onHandleNameChange}/>
                ) : (
                    <Text>{name}</Text>
                )}
              </div>
              <Text strong>Status</Text><br />
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
                  <ProjectStatus status={projectStatus ? projectStatus : 'no status'} />
              )}
          </div>
          <div style={{paddingRight: '24px'}}>
          {hasPrivilege(loggedInUser.privileges, 'user_read') && project && (<>
            <Text strong>Created by</Text><br />
            <Link href={`/user/${project.created_by}`}>{getUserName(project.created_by)}</Link><br />
            <Text strong>Created at</Text><br />
            <Text>{formatDateTimeToYYYYMMDDHHMM(project.created_at)}</Text><br />
            <Text strong>Updated by</Text><br />
            <Link href={`/user/${project.updated_by}`}>{getUserName(project.updated_by)}</Link><br />
            <Text strong>Updated at</Text><br />
            <Text>{formatDateTimeToYYYYMMDDHHMM(project?.updated_at)}</Text><br />
          </>)}
          </div>
          </div>
          <Divider />
          <div style={{display: 'flex', justifyContent: 'flex-end', gap: '8px'}}>
              {editing && hasPrivilege(loggedInUser.privileges, PRIVILEGES.project_sudo) && (
                  <Popconfirm
                      placement="top"
                      title="Are you sure?"
                      description={`Do you want to delete user ${name}`}
                      onConfirm={onClickdeleteProject}
                      icon={<QuestionCircleOutlined twoToneColor="red" />}
                      okText="Yes"
                      cancelText="No"
                  >
                      <Button danger type="link"><DeleteOutlined /></Button>
                  </Popconfirm>
              )}
              <Button type={editing ? "default" : "primary"} onClick={() => setEditing(!editing)}>{editing ? 'Close' : 'Edit'}</Button>
              {editing && (<Button type="primary" onClick={onSaveEdittedProject}>Save</Button>)}
          </div>
          </div>
        ),
        projectEconomics: (
          <Card 
            tabList={economicsTabList}
            activeTabKey={activeEconomicTab}
            style={{borderRadius: 0, border: 'none'}}
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
          style={{ width: '98%', height: 'fit-content', padding: 0}}
          tabList={tabList}
          activeTabKey={activeTab}
          bodyStyle={{padding: '0px'}}
          onTabChange={onHandleChangeActiveTab}
          >
          {contentList[activeTab]}
        </Card>
        </Col>
        <Col span={8}>
        <Card style={{ width: '100%', height: 'fit-content'}}>
          <CreateNote
            type={NOTE_TYPE.project}
            title={noteTitle}
            onTitleChange={onHandleNoteTitleChange}
            note={note}
            onNoteChange={onHandleNoteChange}
            onClearNoteFields={clearNoteFields}
            onSubmit={onSubmitProjectNote}
          />
        {hasPrivilege(loggedInUser.privileges, 'note_read') && projectNotes && (
          <NoteList notes={projectNotes} type={NOTE_TYPE.project} userId={loggedInUser.id} />
        )}
        </Card>
        </Col>
      </Row>
      )
}

export default ProjectDetails;