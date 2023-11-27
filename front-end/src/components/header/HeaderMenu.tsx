/* eslint-disable @typescript-eslint/no-explicit-any */
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { useSelector, useDispatch } from 'react-redux';
import { State } from '../../interfaces/state';
import { LoginOutlined } from '@ant-design/icons';
import { logout } from "../../redux/userDataSlice";
import { clearData } from '../../redux/applicationDataSlice';
import { Button, Space, Typography } from "antd"

const { Text, Title } = Typography;

const headerTitle = (pathName : string) => {
    if (pathName.includes('my-details')) {
        return 'My details'
    }
    if (pathName.includes('/user/')) {
        return 'User details'
    }
    if (pathName.includes('users')) {
        return 'Users'
    }
    if (pathName.includes('sub-projects')) {
        return 'Sub Projects'
    }
    if (pathName.includes('/project/')) {
        return 'Project details'
    }
    if (pathName.includes('projects')) {
        return 'Projects'
    }
    if (pathName.includes('/privilege/')) {
        return 'Privilege details'
    }
    if (pathName.includes('privileges')) {
        return 'Privileges'
    }
    if (pathName.includes('/expense/')) {
        return 'Expense details'
    }
    if (pathName.includes('create-expense')) {
        return 'Create new expense'
    }
    if (pathName.includes('expenses')) {
        return 'Expenses'
    }
    if (pathName.includes('/invoice/')) {
        return 'Invoice details'
    }
    if (pathName.includes('create-invoice')) {
        return 'Create new invoice'
    }
    if (pathName.includes('invoices')) {
        return 'Invoices'
    }
    if (pathName.includes('/income/')) {
        return 'Income details'
    }
    if (pathName.includes('create-income')) {
        return 'Create new income'
    }
    if (pathName.includes('incomes')) {
        return 'Incomes'
    }
    if (pathName.includes('products')) {
        return 'Products'
    }
    if (pathName.includes('/external-company/')) {
        return 'External company details'
    }
    if (pathName.includes('create-external-company')) {
        return 'Create new external company'
    }
    if (pathName.includes('external-companies')) {
        return 'External companies'
    }
    if (pathName.includes('services')) {
        return 'Services overview'
    }

    return 'Manage management'
}

const HeaderMenu = () => {
    const user = useSelector((state: State) => state.user);
    const darkTheme = useSelector((state : State) => state.user.settings.dark_theme);
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

    const backgroundColor = darkTheme ? '#141414' : '#ffffff';
    const borderBottom = darkTheme ? '1px solid #303030' : '1px solid rgba(5, 5, 5, 0.06)';

    return (
        <div style={{ height: '100%', display: 'flex', justifyContent: 'space-between', alignItems: 'center', paddingRight: '8px', paddingLeft: '12px', background: backgroundColor, borderBottom: borderBottom}}>
            <div style={{ height: '100%', display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
                <Title style={{paddingTop: '10px'}} level={5}>{headerTitle(window.location.pathname)}</Title>
            </div>
            <div style={{ height: '100%', display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
                {user.authenticated ? (
                    <Space direction="horizontal">
                        <Text>{user.firstName} {user.lastName}</Text>
                        <Text onClick={() => navigate('/my-details')} underline italic style={{ marginRight: '4px', cursor: 'pointer' }}>{user.email}</Text>
                        <Link to="/login">
                            <Button onClick={() => OnLoginButtonClick(user.authenticated)}>
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