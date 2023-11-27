import { useDispatch, useSelector } from 'react-redux';
import { State } from '../interfaces/state';
import {CheckOutlined, CloseOutlined } from '@ant-design/icons';
import { Typography, Space, Switch, notification, Button, Card } from 'antd';
import { updateUserSettings } from '../api/users/userSettings/updateUserSettings';
import { updateDarkTheme, updateCompactUI } from '../redux/userDataSlice';
import { BlueTag } from '../components/tags/BlueTag';

const { Text, Title } = Typography

const MyPage = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const user = useSelector((state : State) => state.user);
    const isDarkTheme = useSelector((state : State) => state.user.settings.dark_theme);
    const isCompactUI = useSelector((state : State) => state.user.settings.compact_ui);

    if (!user) {
        return <Title level={1}>Something went wrong with fetching your details</Title>;
    }

    const updateDarkThemeToggle = () => {
        dispatch(updateDarkTheme(!isDarkTheme));
    }

    const updateCompactUIToggle = () => {
        dispatch(updateCompactUI(!isCompactUI));
    }

    const saveCosmeticSettings = () => {
        updateUserSettings(user.id, user.id, isDarkTheme, isCompactUI).then(() => {
            api.success({
                message: 'Cosmetic settings updated',
                placement: 'bottom',
                duration: 1.4
            });
        }).catch((error) => {
            api.error({
                message: 'Error updating cosmetic settings',
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    }

    return (
        <div style={{padding: '16px'}}>
            {contextHolder}
            <div style={{width: '100%', display: 'flex', justifyContent: 'flex-start', gap: '24px', marginBottom: '24px'}}>
            <Card style={{width: 'fit-content', height: 'fit-content', minWidth: '400px'}} title="User details">
            <Title level={5}>Name</Title>
            <Text>{user.firstName} {user.lastName}</Text>
            <Title level={5}>Email</Title>
            <Text>{user.email}</Text>
            </Card>
            <Card style={{width: 'fit-content', height: 'fit-content',}} title="Cosmetic settings">
                <div style={{display: 'flex', justifyContent: 'flex-start', alignItems: 'center', marginBottom: '24px', gap: '24px'}}>
                    <Space direction="vertical">
                        <Text strong>Dark mode</Text>
                        <Switch
                            checkedChildren={<CheckOutlined />}
                            unCheckedChildren={<CloseOutlined />}
                            checked={isDarkTheme}
                            onChange={updateDarkThemeToggle}
                        />
                    </Space>
                    <Space direction="vertical">
                    <Text strong>Compact UI mode</Text>
                    <Switch
                            checkedChildren={<CheckOutlined />}
                            unCheckedChildren={<CloseOutlined />}
                            onChange={updateCompactUIToggle}
                            checked={isCompactUI}
                        />
                    </Space>
                </div>
                <div style={{width: '100%', display: 'flex', justifyContent: 'flex-end'}}>
                        <Button type="primary" onClick={saveCosmeticSettings}>Save</Button>
                </div>
            </Card>
            </div>
            <div style={{width: '100%', display: 'flex', justifyContent: 'flex-start', gap: '24px', marginBottom: '24px'}}>
            <Card style={{width: 'fit-content', height: 'fit-content', maxWidth: '800px'}} title="Privileges">
            <div style={{marginBottom: '8px'}}>{user.privileges.map(privilege => <BlueTag label={privilege}/>)}</div>
            </Card>
            </div>
        </div>
    )
}

export default MyPage;