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
  Divider,
} from "antd";
import { useDispatch, useSelector } from "react-redux";
import { getAllProjectNotesByUserId } from '../../api/notes/project/getAllByUserId';
import { useEffect, useState } from "react";
import { updateUserCall } from "../../api/users/update";
import { QuestionCircleOutlined, DeleteOutlined } from "@ant-design/icons";
import { State } from "../../types/state";
import { updateUser, popUser } from "../../redux/applicationDataSlice";
import { deleteUser } from "../../api/users/delete";
import { hasPrivilege } from "../../helpers/hasPrivileges";
import { PRIVILEGES } from "../../enums/privileges";
import { BlueTags } from "../tags/BlueTags";
import { PurpleTags } from "../tags/DefaultTags";
import { SelectOptions } from '../../types/generics';
import { NOTE_TYPE } from "../../enums/notes";
import Notes from "../notes/Notes";
import { getAllExpenseNotesByUserId } from "../../api/notes/expense/getAllByUserId";
import { getAllIncomeNotesByUserId } from "../../api/notes/income/getAllByUserId";
import { getAllExternalCompanyNotesByUserId } from "../../api/notes/externalCompany/getAllByUserId";
import {
  FundProjectionScreenOutlined,
  ProjectOutlined,
  DollarOutlined,
  FundOutlined,
  BankOutlined
} from '@ant-design/icons';
import { getAllSubProjectNotesByUserId } from "../../api/notes/subProject/getAllByUserId";

const { Text, Title } = Typography;

const userNotesTabList = [
  {
    key: 'project',
    label: <FundProjectionScreenOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'subProject',
    label: <ProjectOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'expense',
    label: <DollarOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'income',
    label: <FundOutlined style={{paddingLeft: '12px'}} />,
  },
  {
    key: 'companies',
    label: <BankOutlined style={{paddingLeft: '12px'}} />,
  },
];

const User: React.FC = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [api, contextHolder] = notification.useNotification();

  const [projectNotes, setProjectNotes] = useState([]);
  const [subProjectNotes, setSubProjectNotes] = useState([]);
  const [expenseNotes, setExpenseNotes] = useState([]);
  const [incomeNotes, setIncomeNotes] = useState([]);
  const [externalCompaniesNotes, setExternalCompaniesNotes] = useState([]);
  const [activeNotesTab, setActiveNotesTab] = useState<string>('project');
  
  const { id } = useParams();
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
  const [allProjectsOptions, setAllProjectsOptions] = useState<Array<SelectOptions>>([]);
  const [editing, setEditing] = useState(false);

  useEffect(() => {
    if (user) {
      try {

        setEmail(user.email);
        setFirstName(user.first_name);
        setLastName(user.last_name);
        setPrivileges(user.privileges);
        setProjects(user.projects);

        const userProjects : Array<any> = [];
        allProjects.forEach(project => {
          if (user.projects.includes(project.id)) {
            userProjects.push({ label: project.name, value: project.id})
          }
        });
       
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
  }, [user]);

  useEffect(() => {
    if (loggedInUserId) {
      getAllProjectNotesByUserId(loggedInUserId, userId).then(response => {
        if (response.data?.length) {
          setProjectNotes(response.data);
        } else {
          setProjectNotes([]);
        }
      }).catch((error : any) => {
        console.log('error fetching project notes', error);
      });

      getAllSubProjectNotesByUserId(loggedInUserId, userId).then(response => {
        if (response.data?.length) {
          setSubProjectNotes(response.data);
        } else {
          setSubProjectNotes([]);
        }
      }).catch((error : any) => {
        console.log('error fetching project notes', error);
      });
   
      getAllExpenseNotesByUserId(loggedInUserId, userId).then(response => {
        if (response.data?.length) {
          setExpenseNotes(response.data);
        } else {
          setExpenseNotes([]);
        }
      }).catch((error : any) => {
        console.log('error fetching expense notes', error);
      });

      getAllIncomeNotesByUserId(loggedInUserId, userId).then(response => {
        if (response.data?.length) {
          setIncomeNotes(response.data);
        } else {
          setIncomeNotes([]);
        }
      }).catch((error : any) => {
        console.log('error fetching income notes', error);
      });
      getAllExternalCompanyNotesByUserId(loggedInUserId, userId).then(response => {
        if (response.data?.length) {
          setExternalCompaniesNotes(response.data);
        } else {
          setExternalCompaniesNotes([]);
        }
      }).catch((error : any) => {
        console.log('error fetching income notes', error);
      });
    }
  }, [loggedInUserId, activeNotesTab]);

  const onHandleEmailChange = (event: any) => setEmail(event.target.value);
  const onHandleFirstNameChange = (event: any) => setFirstName(event.target.value);
  const onHandleLastNameChange = (event: any) => setLastName(event.target.value);
  const onHandlePrivilegeChange = (value: any) => setPrivileges(value);
  const onHandleProjectsChange = (value: any) => setProjects(value);
  const onHandleChangeActiveNotesTab = (tab: string) => setActiveNotesTab(tab);

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
          message: response.message,
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
          message: response.message,
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

  const notesContentList: Record<string, React.ReactNode> = {
    project: <Notes notes={projectNotes} type={NOTE_TYPE.project} userId={loggedInUserId} generalized /> ,
    subProject: <Notes notes={subProjectNotes} type={NOTE_TYPE.sub_project} userId={loggedInUserId} generalized /> ,
    expense:  <Notes notes={expenseNotes} type={NOTE_TYPE.expense} userId={loggedInUserId} generalized /> ,
    income: <Notes notes={incomeNotes} type={NOTE_TYPE.income} userId={loggedInUserId} generalized /> ,
    companies: <Notes notes={externalCompaniesNotes} type={NOTE_TYPE.external_company} userId={loggedInUserId} generalized />
  }

  return (
    <>
      {contextHolder}
      <Row>
        <Col span={16}>
          <Card style={{ width: '98%'}}>
          <div style={{display: 'flex', justifyContent: 'flex-start', gap: '20px'}}>
            <div style={{paddingRight: '16px'}}>
            <Title level={4}>User information</Title>
            <Space direction="vertical">
                    <Text strong>First name</Text>
                    <Input
                      value={firstName}
                      disabled={!editing}
                      onChange={onHandleFirstNameChange}
                      style={{ width: 210 }}
                    />
                    <Text strong>Last name</Text>
                    <Input
                      value={lastName}
                      disabled={!editing}
                      onChange={onHandleLastNameChange}
                      style={{ width: 210 }}
                    />
                  <Text strong>Email</Text>
                  <Input
                    value={email}
                    disabled={!editing}
                    onChange={onHandleEmailChange}
                    style={{ width: 320 }}
                  />
            </Space>
            </div>
            <Space direction="vertical" style={{paddingRight: '24px'}}>
              <Title level={4} style={{marginBottom: '0px'}}>User Privileges</Title>
              <Text strong>Privileges connected to the user</Text>
              <Select
                disabled={!editing}
                mode="multiple"
                style={{ width: "100%" }}
                placeholder="Select privilege"
                value={privileges}
                tagRender={BlueTags}
                onChange={onHandlePrivilegeChange}
                options={privilegesOptions}
              />
              <Title level={4} style={{marginTop: '24px', marginBottom: '0px'}}>Projects</Title>
                <Text strong>Projects connected to the user</Text>
                <Select
                  disabled={!editing}
                  mode="multiple"
                  style={{ width: "100%" }}
                  placeholder="Select projects"
                  tagRender={PurpleTags}
                  value={projects}
                  onChange={onHandleProjectsChange}
                  options={allProjectsOptions}
                />
            </Space>
            </div>
            <Divider />
            <div
            style={{
              display: "flex",
              justifyContent: "flex-end",
              gap: "8px",
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
                <Button danger><DeleteOutlined /></Button>
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
          </Card>
        </Col>
        <Col span={8}>
        <Card 
          style={{width: '400px', height: 'fit-content'}}
          tabList={userNotesTabList}
          onTabChange={onHandleChangeActiveNotesTab}
          >
            {notesContentList[activeNotesTab]}
          </Card>
        </Col>
      </Row>
    </>
  );
};

export default User;
