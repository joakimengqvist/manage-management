/* eslint-disable @typescript-eslint/no-explicit-any */
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { deleteProject } from '../../api/projects/delete';
import { Table, Button, Popconfirm, notification } from 'antd';
import { State } from '../../types/state';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined } from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { ProjectStatus, ProjectStatusTypes } from './../tags/ProjectStatus';
import { cardShadow } from '../../enums/styles';
interface Project {
    id: string;
    name: string;
    status: ProjectStatusTypes;
    notes: Array<string>
    created_at: string;
    updated_at: string
    delete_project: any
}

const columns = [
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
        title: '',
        dataIndex: 'operations',
        key: 'operations'
    },
]

const Projects: React.FC = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const projects = useSelector((state : State) => state.application.projects);

    const navigateToProject = (id : string) => navigate(`/project/${id}`);

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
            name: <Button type="link" onClick={() => navigateToProject(project.id.toString())}>{project.name}</Button>,
            status: <ProjectStatus status={project.status} />,
            operations: (<div  style={{display: 'flex', justifyContent: 'flex-end'}}>
                <Button type="link" onClick={() => navigateToProject(project.id.toString())}>Edit</Button>
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
            </div>)
        }
    })


    return  <>{contextHolder}<Table style={{boxShadow: cardShadow, borderTopLeftRadius: '8px', borderTopRightRadius: '8px'}} size="small" bordered columns={columns} dataSource={projectsData} /></>
}

export default Projects;