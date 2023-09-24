/* eslint-disable @typescript-eslint/no-explicit-any */
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { useSelector, useDispatch } from 'react-redux';
import { State } from '../../types/state';
import { LoginOutlined } from '@ant-design/icons';
import { logout } from "../../redux/userDataSlice";
import { clearData, selectProject } from '../../redux/applicationDataSlice';
import { Button, Space, Typography, Select } from "antd"

const { Text } = Typography;

const HeaderMenu: React.FC = () => {
    const user = useSelector((state: State) => state.user)
    const projects = useSelector((state: State) => state.application.projects)
    const dispatch = useDispatch()
    const navigate = useNavigate()

    const OnLoginButtonClick = (isAuth: boolean) => {
        if (isAuth) {
            dispatch(logout())
            dispatch(clearData())
        } else {
            navigate("/login")
        }
    }

    const onSelectProject = (value: any) => {
        const jsonValue = JSON.parse(value)
        dispatch(selectProject(jsonValue))
    }

    const projectOptions = [{label: 'All projects', value: JSON.stringify({id: 0, name: 'All projects'})}]
    projects.forEach(project => {
        projectOptions.push({
            value: JSON.stringify({ id: project.id, name: project.name }),
            label: project.name,
        })
    })

    return (
        <div style={{ height: '100%', display: 'flex', justifyContent: 'space-between', alignItems: 'center', paddingRight: '28px', paddingLeft: '20px', borderBottom: '1px solid #d9d9d9', background: 'white' }}>
            <div style={{ height: '100%', display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
                {user.authenticated && (
                    <Select
                        defaultValue={projectOptions[0].value}
                        style={{ width: 300 }}
                        options={projectOptions}
                        onSelect={onSelectProject}
                    />
                )}
            </div>
            <div style={{ height: '100%', display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
                {user.authenticated ? (
                    <Space direction="horizontal">
                        <Text>{user.firstName} {user.lastName}</Text>
                        <Text onClick={() => navigate('/my-details')} underline italic style={{ marginRight: '4px', cursor: 'pointer' }}>{user.email}</Text>
                        <Link to="/login">
                            <Button type="primary" onClick={() => OnLoginButtonClick(user.authenticated)}>
                                Log out
                            </Button>
                        </Link>
                    </Space>

                ) : (
                    <Button type="primary" onClick={() => OnLoginButtonClick(user.authenticated)}>
                        Log in <LoginOutlined />
                    </Button>
                )}
            </div>
        </div>
    )

}

export default HeaderMenu