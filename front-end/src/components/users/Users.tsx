/* eslint-disable @typescript-eslint/no-explicit-any */
import { useDispatch } from 'react-redux';
import { deleteUser } from '../../api/users/delete';
import { Table, Button, Popconfirm, notification, Typography } from 'antd';
import { popUser } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined, DeleteOutlined, ZoomInOutlined } from '@ant-design/icons';
import { useEffect, useState } from 'react';
import { getAllUsers } from '../../api/users/getAll';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { User } from '../../interfaces/user';
import { useGetLoggedInUser } from '../../hooks';
import LargeLoader from '../loader/LargeLoader';

const { Link } = Typography;

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

const Users = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useGetLoggedInUser();
    const [loading, setLoading] = useState<boolean>(true);
    const [users, setUsers] = useState<Array<User>>([]);

    useEffect(() => {
        if (loggedInUser.id) {
            setLoading(true)
            getAllUsers(loggedInUser.id).then(response => {
                if (response?.error) {
                    api.error({
                        message: `Fetching users failed`,
                        description: response.message,
                        placement: 'bottom',
                        duration: 1.4
                      });
                      setLoading(false)
                      return
                }
                setUsers(response.data)
                setLoading(false)
            }).catch((error) => {
                api.error({
                    message: `Error fetching users`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                })
                setLoading(false)
            });
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [loggedInUser.id])

    const onClickdeleteUser = async (id : string) => {
        await deleteUser(loggedInUser.id, id)
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
                    {hasPrivilege(loggedInUser.privileges, PRIVILEGES.user_sudo) &&
                        <Popconfirm
                            placement="top"
                            title="Are you sure?"
                            description={`Do you want to delete user ${user.first_name}`}
                            onConfirm={() => onClickdeleteUser(user.id)}
                            icon={<QuestionCircleOutlined twoToneColor="red" />}
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

    return (<>
        {contextHolder}
        {loading && <LargeLoader />}
        {!loading && <Table size="small" bordered columns={columns} dataSource={usersData} />}
     </>)
}

export default Users;