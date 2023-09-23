/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useNavigate, useParams } from "react-router-dom";
import {
  Row,
  Col,
  Button,
  Card,
  Space,
  Input,
  Select,
  Typography,
  notification,
  Popconfirm,
} from "antd";
import { useDispatch, useSelector } from "react-redux";
import { useEffect, useState } from "react";
import { updateUserCall } from "../../api/users/update";
import { QuestionCircleOutlined } from "@ant-design/icons";
import { State } from "../../types/state";
import { updateUser, popUser } from "../../redux/applicationDataSlice";
import { deleteUser } from "../../api/users/delete";
import { hasPrivilege } from "../../helpers/hasPrivileges";
import { PRIVILEGES } from "../../enums/privileges";

const { Text, Title } = Typography;

interface Project {
  label : string,
  value : string,
}

const User: React.FC = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [api, contextHolder] = notification.useNotification();

  const { id } = useParams(); // user id
  const userId = id || '';
  const users = useSelector((state: State) => state.application.users);
  const user = users.find((u : any) => u.id === userId);
  const allProjects = useSelector((state: State) => state.application.projects);
  const loggedInUserId = useSelector((state: State) => state.user.id);
  const userPrivileges = useSelector((state: State) => state.user.privileges);
  const allPrivileges = useSelector((state: State) => state.application.privileges);

  const [email, setEmail] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [privilegesOptions, setPrivilegesOptions] = useState<Array<any>>([]);
  const [privileges, setPrivileges] = useState<Array<any>>([]);
  const [projects, setProjects] = useState<Array<any>>([]);
  const [allProjectsOptions, setAllProjectsOptions] = useState<Array<Project>>([]);
  const [editing, setEditing] = useState(false);

  useEffect(() => {
    if (user) {
      try {

        setEmail(user.email);
        setFirstName(user.first_name);
        setLastName(user.last_name);
        setPrivileges(user.privileges);
        setProjects(user.projects);

        const privilegesOptions = allPrivileges.map(privilege => {
          return { label: privilege.name, value: privilege.name }
        });
        setPrivilegesOptions(privilegesOptions);

        const projectsOptions = allProjects.map(project => {
          return { label: project.name, value: project.id }
        });
        setAllProjectsOptions(projectsOptions)

        } catch (error : any) {
          api.error({
            message: `Error fetching user`,
            description: error.toString(),
            placement: "bottom",
            duration: 2,
          });
        }
    }
  }, [users]);

  const onHandleEmailChange = (event: any) => setEmail(event.target.value);
  const onHandleFirstNameChange = (event: any) => setFirstName(event.target.value);
  const onHandleLastNameChange = (event: any) => setLastName(event.target.value);
  const onHandlePrivilegeChange = (value: any) => setPrivileges(value);
  const onHandleProjectsChange = (value: any) => setProjects(value);

  const onSaveEdittedUser = async () => {
    await updateUserCall(
      loggedInUserId,
      userId,
      firstName,
      lastName,
      email,
      privileges,
      projects,
    )
      .then(response => {
        if (response?.error) {
          api.error({
              message: `Updated user failed`,
              description: response.message,
              placement: 'bottom',
              duration: 1.4
            });
          return
        }
        api.info({
          message: `Updated user`,
          description: "Succesfully updated user.",
          placement: "bottom",
          duration: 1.2,
        });
        dispatch(updateUser(response.data));
      })
      .catch((error) => {
        api.error({
          message: `Error updating user`,
          description: error.toString(),
          placement: "bottom",
          duration: 1.4,
        });
      });
  };

  const onClickdeleteUser = async (id: string) => {
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
          description: "Succesfully deleted user.",
          placement: "bottom",
          duration: 1.2,
        });
        dispatch(popUser(id));
        setTimeout(() => {
          navigate("/users");
        }, 1000);
      })
      .catch((error) => {
        api.error({
          message: `Error deleting user`,
          description: error.toString(),
          placement: "bottom",
          duration: 1.4,
        });
      });
  };

  return (
    <Card style={{ padding: "0px" }}>
      {contextHolder}
      <Row>
        <Col style={{ marginRight: "16px" }}>
          <Card>
            <Title level={4}>User information</Title>
            <Space direction="vertical">
                <>
                  <Space direction="vertical">
                    <Text strong>First name</Text>
                    <Input
                      value={firstName}
                      disabled={!editing}
                      onChange={onHandleFirstNameChange}
                      style={{ width: 210 }}
                    />
                  </Space>
                  <Space direction="vertical">
                    <Text strong>Last name</Text>
                    <Input
                      value={lastName}
                      disabled={!editing}
                      onChange={onHandleLastNameChange}
                      style={{ width: 210 }}
                    />
                  </Space>
                  <Text strong>Email</Text>
                  <Input
                    value={email}
                    disabled={!editing}
                    onChange={onHandleEmailChange}
                    style={{ width: 320 }}
                  />
                </>
            </Space>
          </Card>
        </Col>
        <Col>
          <Card style={{ width: "460px" }}>
            <Space direction="vertical" style={{ width: "100%" }}>
              <Title level={4} style={{marginBottom: '0px'}}>User Privileges</Title>
              <Text strong>Privileges connected to the user</Text>
              <Select
                disabled={!editing}
                mode="multiple"
                style={{ width: "100%" }}
                placeholder="Select privilege"
                value={privileges}
                onChange={onHandlePrivilegeChange}
                options={privilegesOptions}
              />
            </Space>
            <Space direction="vertical" style={{ width: "100%" }}>
              <Title level={4} style={{marginTop: '24px', marginBottom: '0px'}}>Projects</Title>
              <Text strong>Projects connected to the user</Text>
              <Select
                disabled={!editing}
                mode="multiple"
                style={{ width: "100%" }}
                placeholder="Select privilege"
                value={projects}
                onChange={onHandleProjectsChange}
                options={allProjectsOptions}
              />
            </Space>
          </Card>
        </Col>
      </Row>
      <Row>
        <Col span={16}>
          <div
            style={{
              display: "flex",
              justifyContent: "flex-start",
              gap: "12px",
              paddingTop: "24px",
            }}
          >
            {editing && hasPrivilege(userPrivileges, PRIVILEGES.user_sudo) && (
              <Popconfirm
                placement="top"
                title="Are you sure?"
                description={`Do you want to delete user ${firstName}`}
                onConfirm={() => onClickdeleteUser(userId)}
                icon={<QuestionCircleOutlined style={{ color: "red" }} />}
                okText="Yes"
                cancelText="No"
              >
                <Button danger>Delete</Button>
              </Popconfirm>
            )}
            <div
              style={{
                display: "flex",
                justifyContent: "flex-start",
                gap: "8px",
              }}
            >
              <Button
                type={editing ? "default" : "primary"}
                onClick={() => setEditing(!editing)}
              >
                {editing ? "Close" : "Edit"}
              </Button>
              {editing && (
                <Button type="primary" onClick={onSaveEdittedUser}>
                  Save
                </Button>
              )}
            </div>
          </div>
        </Col>
      </Row>
    </Card>
  );
};

export default User;
