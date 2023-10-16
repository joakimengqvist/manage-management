/* eslint-disable @typescript-eslint/no-explicit-any */
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { Table, Button, Popconfirm, notification, Modal, Select, Typography } from 'antd';
import { State } from '../../types/state';
import { popProject } from '../../redux/applicationDataSlice';
import { 
    QuestionCircleOutlined, 
    DeleteOutlined, 
    ZoomInOutlined,
    SettingOutlined,
} from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { ProjectStatus } from './../tags/ProjectStatus';
import { SubProject } from '../../types/subProject';
import { deleteSubProject } from '../../api/subProjects/delete';
import Priority from '../renderHelpers/Priority';
import EstimatedDuration from '../renderHelpers/EstimatedDuration';
import { useState } from 'react';
import { AddProjectSubProjectConnection } from '../../api/subProjects/specialActions/addProjectSubProjectConnection';

const { Text } = Typography;

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
    /*
    {
        title: 'Project',
        dataIndex: 'project',
        key: 'project'
    },
    */
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

    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isModalLoading, setIsModalLoading] = useState(false);
    const [modalSelectedSubProjectId, setModalSubSelectedProjectId] = useState('');
    const [modalSelectedProjectId, setModalSelectedProjectId] = useState('');

    const getSubProjectName = (id : string) => subProjects.find(subProject => subProject.id === id)?.name;

    const openModal = (subProjectId : string) => {
        setModalSubSelectedProjectId(subProjectId);
        setIsModalOpen(true);
    }

    const onModalConfirm = () => {
        setIsModalLoading(true);
        AddProjectSubProjectConnection(userId, modalSelectedSubProjectId, modalSelectedProjectId).then(response => {
            if (response?.error) {
                api.error({
                    message: `Error adding project to sub project`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                });
                return;
            }
            api.info({
                message: `added project to sub project`,
                description: 'Succesfully added project to sub project',
                placement: 'bottom',
                duration: 1.4
            });
            setIsModalOpen(false);
        }).catch(error => {
            api.error({
                message: `Error adding project to sub project`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
            setIsModalOpen(false);
        });
        setIsModalLoading(false);
    }

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

    const projectOptions = projects.map(project => {
        return { label: project.name, value: project.id}
      }
    );

    const subProjectsData: Array<any> = subProjects && subProjects.map((subProject : SubProject) => ({                 
            name: <Button type="link" onClick={() => navigate(`/sub-project/${subProject.id}`)}>{subProject.name}</Button>,
            status: <ProjectStatus status={subProject.status} />,
            priority: <Priority priority={subProject.priority} />,
            estimated_duration: <EstimatedDuration duration={subProject.estimated_duration} />,
            /* project: <Link to={`/project/${subProject.project_id}`}>{getProjectName(subProject.project_id)}</Link>, */
            operations: (
                <div  style={{display: 'flex', justifyContent: 'flex-end'}}>
                    <Button style={{padding:'8px'}} type="link" onClick={() => openModal(subProject.id)}><SettingOutlined /></Button>
                    <Button style={{padding:'8px'}} type="link" onClick={() => navigate(`/sub-project/${subProject.id}`)}><ZoomInOutlined /></Button>
                    {hasPrivilege(userPrivileges, PRIVILEGES.sub_project_sudo) &&
                    <Popconfirm
                        placement="top"
                        title="Are you sure?"
                        description={`Do you want to delete project ${subProject.name}`}
                        onConfirm={() => onClickdeleteProject(subProject.id)}
                        icon={<QuestionCircleOutlined style={{ color: 'red' }} />}
                        okText="Yes"
                        cancelText="No"
                    >
                        <Button danger style={{padding:'8px'}} type="link"><DeleteOutlined /></Button>
                    </Popconfirm>
                    }
                </div>
            )
        }));

    if (!subProjectsData) return null;


    return  <>
    {contextHolder}
        <Table size="small" bordered columns={columns} dataSource={subProjectsData} />
        <Modal
            title={`Add project to ${getSubProjectName(modalSelectedSubProjectId)}`}
            open={isModalOpen}
            onOk={onModalConfirm}
            confirmLoading={isModalLoading}
            onCancel={() => setIsModalOpen(false)}
        > 
            <Text style={{paddingBottom: '16px'}}>What project do you want to add {getSubProjectName(modalSelectedSubProjectId)} to?</Text>
            <Select
            style={{width: '100%', marginTop: '16px', marginBottom: '8px'}}
            options={projectOptions}
            onChange={(value : string) => setModalSelectedProjectId(value)}
            value={modalSelectedProjectId}
            />
    </Modal>
    </>
}

export default SubProjects;