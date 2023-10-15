/* eslint-disable @typescript-eslint/no-explicit-any */
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { deleteProject } from '../../api/projects/delete';
import { Table, Button, Popconfirm, notification } from 'antd';
import { State } from '../../types/state';
import { Project } from '../../types/project';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined } from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { ProjectStatus } from './../tags/ProjectStatus';
import { SubProject } from '../../types/subProject';
import { useMemo } from 'react';
import Priority from '../renderHelpers/RenderPriority';
import RenderEstimatedDuration from '../renderHelpers/RenderEstimatedDuration';

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
                        message: `Deleted user project`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                      });
                    return
                }
                api.info({
                    message: `Deleted project`,
                    description: 'Succesfully deleted project.',
                    placement: 'bottom',
                    duration: 1.4
                  });
                dispatch(popProject(id))
            })
            .catch(error => {
                api.error({
                    message: `Error deleting privilege`,
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
                <Button type="link" onClick={() => navigate(`/project/${project.id}`)}>Edit</Button>
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
                    <Button danger type="link">Delete</Button>
                </Popconfirm>
                }
            </div>),
        }
    })

    

    const expandableProps = useMemo(() => {
        const subProjectData = (id : string) => {
            if (subProjects.length === 0) return [];
            const subProjectsData = subProjects.filter((subProject : SubProject) => subProject.project_id === id);
            const subProjectsDataReturned = subProjectsData.map((subProject : any) => {
                return {
                    name: <Button type="link" onClick={() => navigate(`/sub-project/${subProject.id}`)}>{subProject.name}</Button>,
                    status: <ProjectStatus status={subProject.status} />,
                    priority: <Priority priority={subProject.priority} />,
                    estimated_duration: <RenderEstimatedDuration duration={subProject.estimated_duration} />,
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
                                <Table.Summary.Cell index={1} colSpan={4}>
                                <div style={{display: 'flex', justifyContent: 'flex-end', gap: '12px', padding: '12px 0px 8px 0px', borderTop: '1px solid #f0f0f0'}}>
                                    <Button onClick={() => navigate(`/create-sub-project/project-id/${record.name.props.id}`)}>New expense</Button>
                                    <Button onClick={() => navigate(`/create-sub-project/project-id/${record.name.props.id}`)}>New income</Button>
                                    <Button onClick={() => navigate(`/create-sub-project/project-id/${record.name.props.id}`)}>New sub project</Button>
                                </div>
                                </Table.Summary.Cell>
                            </Table.Summary.Row>
                        </Table.Summary>
                    )}
                />
            ),

        };
      }, [navigate, subProjects]);


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