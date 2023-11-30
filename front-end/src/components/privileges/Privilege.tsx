/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useNavigate, useParams } from 'react-router-dom'
import { Button, Card, Space, Input, Typography, notification, Popconfirm } from 'antd';
import { DeleteOutlined } from '@ant-design/icons';
import { useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';
import { updatePrivilege } from '../../api/privileges/update';
import { QuestionCircleOutlined } from "@ant-design/icons";
import { deletePrivilege } from '../../api/privileges/delete';
import { popPrivilege } from '../../redux/applicationDataSlice';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { getPrivilegeById } from '../../api';
import { useGetLoggedInUser } from '../../hooks';

const { Text, Title } = Typography;

const Privilege = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useGetLoggedInUser();
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [editing, setEditing] = useState(false);
    const { id } =  useParams(); 
    const privilegeId = id || '';

    useEffect(() => {
        getPrivilegeById(loggedInUser.id, privilegeId).then(response => {
            if (response?.error) {
                api.error({
                    message: `Error fetching privilege`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                  });
                  return
              }
            setName(response.data.name);
            setDescription(response.data.description);
        })
    }, [privilegeId]);

    const onHandleNameChange = (event : any) => setName(event.target.value);
    const onHandleDescriptionChange = (event : any) => setDescription(event.target.value);

    const onSaveEdittedPrivilege = async () => {
        await updatePrivilege(loggedInUser.id, privilegeId, name, description)
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
                setEditing(false)
                api.info({
                    message: response.message,
                    placement: 'bottom',
                    duration: 1.4
                });
            })
            .catch(error => {
                api.error({
                    message: `Error Updating privilege`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                });
            })
    };

    const onClickdeletePrivilege = async () => {
        await deletePrivilege(loggedInUser.id, privilegeId)
          .then(response => {
            if (response?.error) {
              api.error({
                  message: `Delete privilege failed`,
                  description: response.message,
                  placement: 'bottom',
                  duration: 1.4
                });
              return
            }
            api.info({
              message: response.message,
              placement: "bottom",
              duration: 1.2,
            });
            dispatch(popPrivilege(privilegeId));
            setTimeout(() => {
              navigate("/privileges");
            }, 1000);
          })
          .catch((error) => {
            api.error({
              message: `Error deleting project`,
              description: error.toString(),
              placement: "bottom",
              duration: 1.4,
            });
          });
      };

    return (
        <Card style={{maxWidth: '400px'}}>
            {contextHolder}
            <Title level={4}>Privilege information</Title>
            <Space direction="vertical" style={{width: '100%'}}>
                <Text strong>Privilege name</Text>
                {editing ? (
                    <Input value={name} onChange={onHandleNameChange}/>
                ) : (
                    <Text>{name}</Text>
                )}
                <Text strong>Privilege desription</Text>
                {editing ? (
                    <Input value={description} onChange={onHandleDescriptionChange}/>
                ) : (
                    <Text>{description}</Text>
                )}
                <div style={{display: 'flex', justifyContent: editing ? 'space-between' : 'flex-end', paddingTop: '18px'}}>
                    {editing && hasPrivilege(loggedInUser.privileges, PRIVILEGES.privilege_sudo) && (
                        <Popconfirm
                            placement="top"
                            title="Are you sure?"
                            description={`Do you want to delete privilege ${name}`}
                            onConfirm={onClickdeletePrivilege}
                            icon={<QuestionCircleOutlined twoToneColor="red" />}
                            okText="Yes"
                            cancelText="No"
                        >
                            <Button danger><DeleteOutlined /></Button>
                        </Popconfirm>
                    )}
                    <div style={{display: 'flex', justifyContent: 'flex-end', gap: '8px'}}>
                    <Button type={editing ? "default" : "primary"} onClick={() => setEditing(!editing)}>{editing ? 'Close' : 'Edit'}</Button>
                    {editing && (<>
                        <Button type="primary" onClick={onSaveEdittedPrivilege}>Save</Button>
                        
                    </>)}
                    </div>
                </div>
            </Space>
        </Card>
    )

}

export default Privilege;