/* eslint-disable @typescript-eslint/no-explicit-any */
import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { Table, Button, Popconfirm, notification } from 'antd';
import { State } from '../../types/state';
import { popProject } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined } from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { ProjectStatus } from './../tags/ProjectStatus';
import { SubProject } from '../../types/subProject';
import { deleteSubProject } from '../../api/subProjects/delete';
import Priority from '../renderHelpers/RenderPriority';
import RenderEstimatedDuration from '../renderHelpers/RenderEstimatedDuration';

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
        title: 'Priority',
        dataIndex: 'priority',
        key: 'priority',
    },
    {
        estimated_duration: 'Estimated duration',
        dataIndex: 'estimated_duration',
        key: 'estimated_duration'
    },
    {
        title: 'Project',
        dataIndex: 'project',
        key: 'project'
    },
    {
        title: '',
        dataIndex: 'operations',
        key: 'operations'
    },
]

const SubProjects: React.FC = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const projects = useSelector((state : State) => state.application.projects);
    const subProjects = useSelector((state : State) => state.application.subProjects);

    const getProjectName = (id : string) => projects.find(project => project.id === id)?.name;

    const onClickdeleteProject = async (id : string) => {
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
                    message: `Deleted subProject`,
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

    const subProjectsData: Array<any> = subProjects.map((subProject : SubProject) => {
        return {                    
            name: <Button type="link" onClick={() => navigate(`/sub-project/${subProject.id}`)}>{subProject.name}</Button>,
            status: <ProjectStatus status={subProject.status} />,
            priority: <Priority priority={subProject.priority} />,
            estimated_duration: <RenderEstimatedDuration duration={subProject.estimated_duration} />,
            project: <Link to={`/project/${subProject.project_id}`}>{getProjectName(subProject.project_id)}</Link>,
            operations: (<div  style={{display: 'flex', justifyContent: 'flex-end'}}>
                <Button type="link" onClick={() => navigate(`/sub-project/${subProject.id}`)}>Edit</Button>
                {hasPrivilege(userPrivileges, PRIVILEGES.project_sudo) &&
                <Popconfirm
                    placement="top"
                    title="Are you sure?"
                    description={`Do you want to delete project ${subProject.name}`}
                    onConfirm={() => onClickdeleteProject(subProject.id)}
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


    return  <>{contextHolder}<Table size="small" bordered columns={columns} dataSource={subProjectsData} /></>
}

export default SubProjects;