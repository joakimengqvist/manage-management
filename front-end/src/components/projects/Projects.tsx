/* eslint-disable @typescript-eslint/no-explicit-any */
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { deleteProject } from '../../api/projects/delete';
import { Table, Button, Popconfirm, notification } from 'antd';
import { State } from '../../types/state';
import { Project } from '../../types/project';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined, DeleteOutlined, ZoomInOutlined } from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { ProjectStatus } from './../tags/ProjectStatus';
import { SubProject } from '../../types/subProject';
import { useMemo } from 'react';
import Priority from '../renderHelpers/Priority';
import EstimatedDuration from '../renderHelpers/EstimatedDuration';
import { deleteSubProject } from '../../api/subProjects/delete';

const projectColumns = [
    {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
    },
    {
        title: 'Status',
        dataIndex: 'status',
        key: 'status'
    },
    {
        title: '',
        dataIndex: 'operations',
        key: 'operations'
    },
];

const subProjectColumns = [
    {
        title: 'Name',
        dataIndex: 'name',
        key: 'name'
    },
    {
        title: 'Status',
        dataIndex: 'status',
        key: 'status'
    },
    {
        title: 'Priority',
        dataIndex: 'priority',
        key: 'priority'  
    },
    {
        title: 'Estimated duration',
        dataIndex: 'estimated_duration',
        key: 'estimated_duration'  
    },
    {
        operations: 'Operations',
        dataIndex: 'operations',
        key: 'operations'
    }
];

const Projects: React.FC = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const projects = useSelector((state : State) => state.application.projects);
    const subProjects = useSelector((state : State) => state.application.subProjects);

    const onClickdeleteProject = async (id : string) => {
        await deleteProject(userId, id)
            .then(response => {
                if (response?.error) {
                    api.error({
                        message: `Deleted project project`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                      });
                    return
                }
                api.info({
                    message: `Deleted project`,
                    description: 'Succesfully deleted sub project.',
                    placement: 'bottom',
                    duration: 1.4
                  });
                dispatch(popProject(id))
            })
            .catch(error => {
                api.error({
                    message: `Error deleting project`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                });
            })
    };

    const projectsData: Array<any> = projects.map((project : Project) => {
        return {      
            id: project.id,              
            name: <Button type="link" id={project.id} onClick={() => navigate(`/project/${project.id}`)}>{project.name}</Button>,
            status: <ProjectStatus status={project.status} />,
            operations: (<div  style={{display: 'flex', justifyContent: 'flex-end'}}>
                <Button type="link" onClick={() => navigate(`/project/${project.id}`)}><ZoomInOutlined /></Button>
                {hasPrivilege(userPrivileges, PRIVILEGES.project_sudo) &&
                <Popconfirm
                    placement="top"
                    title="Are you sure?"
                    description={`Do you want to delete project ${project.name}`}
                    onConfirm={() => onClickdeleteProject(project.id)}
                    icon={<QuestionCircleOutlined style={{ color: 'red' }} />}
                    okText="Yes"
                    cancelText="No"
                >
                    <Button danger type="link"><DeleteOutlined /></Button>
                </Popconfirm>
                }
            </div>),
        }
    })

    const expandableProps = useMemo(() => {

        const onClickdeleteSubProject = async (id : string) => {
            await deleteSubProject(userId, id)
                .then(response => {
                    if (response?.error) {
                        api.error({
                            message: `Deleted sub project`,
                            description: response.message,
                            placement: 'bottom',
                            duration: 1.4
                          });
                        return
                    }
                    api.info({
                        message: `Deleted project`,
                        description: 'Succesfully deleted sub project.',
                        placement: 'bottom',
                        duration: 1.4
                      });
                    dispatch(popProject(id))
                })
                .catch(error => {
                    api.error({
                        message: `Error deleting sub project`,
                        description: error.toString(),
                        placement: 'bottom',
                        duration: 1.4
                    });
                })
        };

        const subProjectData = (projectId : string) => {
            if (!subProjects || subProjects.length === 0) return [];

            const subProjectForProject = subProjects.filter((subProject : SubProject) => subProject.projects.includes(projectId));

            const subProjectsDataReturned = subProjectForProject.map((subProject : any) => {
                return {
                    name: <Button type="link" onClick={() => navigate(`/sub-project/${subProject.id}`)}>{subProject.name}</Button>,
                    status: <ProjectStatus status={subProject.status} />,
                    priority: <Priority priority={subProject.priority} />,
                    estimated_duration: <EstimatedDuration duration={subProject.estimated_duration} />,
                    operations: (<div  style={{display: 'flex', justifyContent: 'flex-end'}}>
                        <Button type="link" onClick={() => navigate(`/sub-project/${subProject.id}`)}><ZoomInOutlined /></Button>
                        {hasPrivilege(userPrivileges, PRIVILEGES.project_sudo) &&
                            <Popconfirm
                                placement="top"
                                title="Are you sure?"
                                description={`Do you want to delete sub project ${subProject.name}`}
                                onConfirm={() => onClickdeleteSubProject(subProject.id)}
                                icon={<QuestionCircleOutlined style={{ color: 'red' }} />}
                                okText="Yes"
                                cancelText="No"
                            >
                                <Button danger type="link"><DeleteOutlined /></Button>
                            </Popconfirm>
                        }
                    </div>),
                }
            });
            return subProjectsDataReturned
        }

        return {
            columnWidth: 48,
            expandedRowRender: (record : any) => (
                <Table 
                    size="small" 
                    bordered 
                    columns={subProjectColumns} 
                    dataSource={subProjectData(record.name.props.id)} 
                    summary={() => (
                        <Table.Summary>
                            <Table.Summary.Row>
                                <Table.Summary.Cell index={1} colSpan={5}>
                                <div style={{display: 'flex', justifyContent: 'flex-end', gap: '12px', padding: '12px 0px 8px 0px', borderTop: '1px solid #f0f0f0'}}>
                                    <Button onClick={() => navigate(`/create-expense/sub-project-id/${record.name.props.id}`)}>New expense</Button>
                                    <Button onClick={() => navigate(`/create-income/sub-project-id/${record.name.props.id}`)}>New income</Button>
                                    <Button onClick={() => navigate(`/create-sub-project/project-id/${record.name.props.id}`)}>New sub project</Button>
                                </div>
                                </Table.Summary.Cell>
                            </Table.Summary.Row>
                        </Table.Summary>
                    )}
                />
            ),
        };
      // eslint-disable-next-line react-hooks/exhaustive-deps
      }, [subProjects, userId, userPrivileges]);

    return (<>
        {contextHolder}
        <Table 
            size="small" 
            rowKey="id"
            bordered 
            columns={projectColumns} 
            dataSource={projectsData} 
            expandable={expandableProps}
        />
     </>)
}

export default Projects;