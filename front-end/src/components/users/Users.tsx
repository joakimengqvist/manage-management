/* eslint-disable @typescript-eslint/no-explicit-any */
import { useDispatch, useSelector } from 'react-redux';
import { deleteUser } from '../../api/users/delete';
import { Table, Button, Popconfirm, notification, Typography } from 'antd';
import { State } from '../../types/state';
import { fetchUsers, popUser } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined, DeleteOutlined, ZoomInOutlined } from '@ant-design/icons';
import { useEffect } from 'react';
import { getAllUsers } from '../../api/users/getAll';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';

const { Link } = Typography;
interface User {
    id: string;
    email: string;
    first_name: string;
    last_name: string;
    created_at: string;
    updated_at: string
}

const columns = [
    {
        title: 'First name',
        dataIndex: 'first_name',
        key: 'first_name'
    },
    {
        title: 'Last name',
        dataIndex: 'last_name',
        key: 'last_name'
    },
    {
        title: 'Email',
        dataIndex: 'email',
        key: 'email'
    },
    {
        title: '',
        dataIndex: 'operations',
        key: 'operations'
    },
]

const Users: React.FC = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const users = useSelector((state : State) => state.application.users);
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    useEffect(() => {
        if (loggedInUserId) {
            getAllUsers(loggedInUserId).then(response => dispatch(fetchUsers(response.data))).catch(() => {});
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [loggedInUserId])

    const onClickdeleteUser = async (id : string) => {
        await deleteUser(loggedInUserId, id)
            .then(response => {
                if (response?.error) {
                    api.error({
                        message: `Deleted user failed`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                      });
                      return
                }
                api.info({
                    message: response.message,
                    placement: 'bottom',
                    duration: 1.4
                  });
                dispatch(popUser(id))
            })
            .catch(error => {
                api.error({
                    message: `Error deleting user`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                });
            })
    };

    const usersData: Array<any> = users.map((user : User) => {
        return {                    
            first_name: <Link href={ `/user/${user.id}`}>{user.first_name}</Link>,
            last_name: user.last_name,
            email: user.email,
            operations: (
                <div style={{display: 'flex', justifyContent: 'flex-end'}}>
                    <Link style={{padding: '5px'}} href={ `/user/${user.id}`}><ZoomInOutlined /></Link>
                    {hasPrivilege(userPrivileges, PRIVILEGES.user_sudo) &&
                        <Popconfirm
                            placement="top"
                            title="Are you sure?"
                            description={`Do you want to delete user ${user.first_name}`}
                            onConfirm={() => onClickdeleteUser(user.id)}
                            icon={<QuestionCircleOutlined style={{ color: 'red' }} />}
                            okText="Yes"
                            cancelText="No"
                        >
                            <Button style={{ padding: '4px' }} danger type="link"><DeleteOutlined /></Button>
                        </Popconfirm>
                    }
                </div>
            )
        }
    });

    return  <>{contextHolder}<Table size="small" bordered columns={columns} dataSource={usersData} /></>
}

export default Users;