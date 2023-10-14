/* eslint-disable @typescript-eslint/no-explicit-any */
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { deletePrivilege } from '../../api/privileges/delete';
import { Table, Button, Popconfirm, notification, Tag } from 'antd';
import { State } from '../../types/state';
import { popPrivilege } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined } from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';

const columns = [
    {
        title: 'Name',
        dataIndex: 'name',
        key: 'name'
    },
    {
        title: 'Description',
        dataIndex: 'description',
        key: 'description'
    },
    {
        title: '',
        dataIndex: 'operations',
        key: 'operations'
    },
]

const Privileges: React.FC = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const privileges = useSelector((state : State) => state.application.privileges);

    const navigateToPrivilege = (id : string) => navigate(`/privilege/${id}`);

    const onClickdeletePrivilege = async (id : string) => {
        await deletePrivilege(userId, id)
            .then(response => {
                if (response?.error) {
                    api.error({
                        message: `Updated privilege failed`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                      });
                    return
                }
                api.info({
                    message: `Deleted privilege`,
                    description: 'Succesfully deleted privilege.',
                    placement: 'bottom',
                    duration: 1.4
                });
                dispatch(popPrivilege(id))
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

    const privilegesData: Array<any> = privileges.map((privilege : any) => {
        return {                    
            name: <Tag color="blue" style={{cursor: 'pointer'}} onClick={() => navigateToPrivilege(privilege.id)}>{privilege.name}</Tag>,
            description: privilege.description,
            operations: (<div style={{display: 'flex', justifyContent: 'flex-end'}}>
                <Button type="link" onClick={() => navigateToPrivilege(privilege.id)}>Details</Button>
                {hasPrivilege(userPrivileges, PRIVILEGES.privilege_sudo) &&
                    <Popconfirm
                        placement="top"
                        title="Are you sure?"
                        description={`Do you want to delete privilege ${privilege.name}`}
                        onConfirm={() => onClickdeletePrivilege(privilege.id)}
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

    return  <>{contextHolder}<Table size="small" bordered columns={columns} dataSource={privilegesData} /></>
}

export default Privileges;