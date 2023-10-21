/* eslint-disable react-hooks/exhaustive-deps */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { useNavigate, useParams } from 'react-router-dom'
import { Button, Card, Space, Input, Typography, notification, Popconfirm } from 'antd';
import { DeleteOutlined } from '@ant-design/icons';
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { updatePrivilege } from '../../api/privileges/update';
import { State } from '../../types/state';
import { QuestionCircleOutlined } from "@ant-design/icons";
import { deletePrivilege } from '../../api/privileges/delete';
import { popPrivilege } from '../../redux/applicationDataSlice';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';

const { Text, Title } = Typography;

const Privilege: React.FC = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUserId = useSelector((state : State) => state.user.id);
    const userPrivileges = useSelector((state : State) => state.user.privileges);
    const privileges = useSelector((state : State) => state.application.privileges);
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [editing, setEditing] = useState(false);
    const { id } =  useParams(); 
    const privilegeId = id || '';

    useEffect(() => {
        const privilege = privileges.find((p : any) => p.id === id);
        if (privilege) {
          try {
            setName(privilege.name);
            setDescription(privilege.description);
            } catch (error : any) {
              api.error({
                message: `Error fetching privilege`,
                description: error.toString(),
                placement: "bottom",
                duration: 1.4,
              });
            }
        }
      }, [privileges]);

    const onHandleNameChange = (event : any) => setName(event.target.value);
    const onHandleDescriptionChange = (event : any) => setDescription(event.target.value);

    const onSaveEdittedPrivilege = async () => {
        await updatePrivilege(loggedInUserId, privilegeId, name, description)
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
        await deletePrivilege(loggedInUserId, privilegeId)
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
                    {editing && hasPrivilege(userPrivileges, PRIVILEGES.privilege_sudo) && (
                        <Popconfirm
                            placement="top"
                            title="Are you sure?"
                            description={`Do you want to delete privilege ${name}`}
                            onConfirm={onClickdeletePrivilege}
                            icon={<QuestionCircleOutlined style={{ color: "red" }} />}
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