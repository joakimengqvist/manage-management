/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { deletePrivilege } from '../../api/privileges/delete';
import { Table, Button, Popconfirm, notification, Typography } from 'antd';
import { popPrivilege } from '../../redux/applicationDataSlice';
import { QuestionCircleOutlined, DeleteOutlined, ZoomInOutlined } from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { BlueTag } from '../tags/BlueTag';
import { useEffect } from 'react';
import { getAllPrivileges } from '../../api';
import { Privilege } from '../../interfaces';
import { useGetLoggedInUser } from '../../hooks';

const { Link } = Typography;

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

const Privileges = () => {
    const dispatch = useDispatch();
    const loggedInUser = useGetLoggedInUser();
    const [api, contextHolder] = notification.useNotification();
    const [privileges, setPrivileges] = useState<Array<Privilege>>([]);

    useEffect(() => {
        getAllPrivileges(loggedInUser.id).then(response => {
            if (response?.error) {
                api.error({
                    message: `Error fetching privileges`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                  });
                  return
              }
            setPrivileges(response.data)
        })
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])

    const onClickdeletePrivilege = async (id : string) => {
        await deletePrivilege(loggedInUser.id, id)
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
                    message: response.message,
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

    const privilegesData: Array<any> = privileges && privileges.length ? privileges.map((privilege : any) => {
        return {                    
            name: <Link style={{cursor: 'pointer'}} href={`/privilege/${privilege.id}`}><BlueTag label={privilege.name} /></Link>,
            description: privilege.description,
            operations: (<div style={{display: 'flex', justifyContent: 'flex-end'}}>
                <Link style={{padding: '5px'}} href={`/privilege/${privilege.id}`}><ZoomInOutlined /></Link>
                {hasPrivilege(loggedInUser.privileges, PRIVILEGES.privilege_sudo) &&
                    <Popconfirm
                        placement="top"
                        title="Are you sure?"
                        description={`Do you want to delete privilege ${privilege.name}`}
                        onConfirm={() => onClickdeletePrivilege(privilege.id)}
                        icon={<QuestionCircleOutlined twoToneColor="red" />}
                        okText="Yes"
                        cancelText="No"
                    >
                        <Button style={{ padding: '4px' }} danger type="link"><DeleteOutlined /></Button>
                    </Popconfirm>
                }
            </div>)
        }
    }) : [];

    return  <>{contextHolder}<Table size="small" bordered columns={columns} dataSource={privilegesData} /></>
}

export default Privileges;