/* eslint-disable @typescript-eslint/no-explicit-any */
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { deleteUser } from '../../api/users/delete';
import { Table, Button, Popconfirm, notification } from 'antd';
import { State } from '../../types/state';
import { fetchUsers, popUser } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined } from '@ant-design/icons';
import { useEffect } from 'react';
import { getAllUsers } from '../../api/users/getAll';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';

interface User {
    id: number;
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
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const users = useSelector((state : State) => state.application.users);
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    useEffect(() => {
        if (loggedInUserId) {
            getAllUsers(loggedInUserId).then(response => dispatch(fetchUsers(response))).catch(() => {});
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [loggedInUserId])

    const navigateToUser = (id : string) => navigate(`/user/${id}`);

    const onClickdeleteUser = async (id : number) => {
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
                    message: `Deleted user`,
                    description: 'Succesfully deleted user.',
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
            first_name: <Button type="link" onClick={() => navigateToUser(user.id.toString())}>{user.first_name}</Button>,
            last_name: user.last_name,
            email: user.email,
            operations: (<>
                <Button type="link" onClick={() => navigateToUser(user.id.toString())}>Edit</Button>
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
                        <Button danger type="link">Delete</Button>
                    </Popconfirm>
                }
            </>)
        }
    });

    return  <>{contextHolder}<Table size="small" bordered columns={columns} dataSource={usersData} /></>
}

export default Users;