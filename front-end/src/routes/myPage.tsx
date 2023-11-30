import { useDispatch } from 'react-redux';
import {CheckOutlined, CloseOutlined } from '@ant-design/icons';
import { Typography, Space, Switch, notification, Button, Card } from 'antd';
import { updateUserSettings } from '../api/users/userSettings/updateUserSettings';
import { updateDarkTheme, updateCompactUI } from '../redux/userDataSlice';
import { BlueTag } from '../components/tags/BlueTag';
import { useGetDarkThemeEnabled, useGetLoggedInUser, useGetCompactUIEnabled } from '../hooks';

const { Text, Title } = Typography

const MyPage = () => {
    const dispatch = useDispatch();
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useGetLoggedInUser();
    const darkTheme = useGetDarkThemeEnabled();
    const compactUi = useGetCompactUIEnabled();

    if (!loggedInUser) {
        return <Title level={1}>Something went wrong with fetching your details</Title>;
    }

    const updateDarkThemeToggle = () => {
        dispatch(updateDarkTheme(!darkTheme));
    }

    const updateCompactUIToggle = () => {
        dispatch(updateCompactUI(!compactUi));
    }

    const saveCosmeticSettings = () => {
        updateUserSettings(
            loggedInUser.id,
            loggedInUser.id,
            darkTheme,
            compactUi
        ).then(() => {
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
            <Text>{loggedInUser.firstName} {loggedInUser.lastName}</Text>
            <Title level={5}>Email</Title>
            <Text>{loggedInUser.email}</Text>
            </Card>
            <Card style={{width: 'fit-content', height: 'fit-content',}} title="Cosmetic settings">
                <div style={{display: 'flex', justifyContent: 'flex-start', alignItems: 'center', marginBottom: '24px', gap: '24px'}}>
                    <Space direction="vertical">
                        <Text strong>Dark mode</Text>
                        <Switch
                            checkedChildren={<CheckOutlined />}
                            unCheckedChildren={<CloseOutlined />}
                            checked={darkTheme}
                            onChange={updateDarkThemeToggle}
                        />
                    </Space>
                    <Space direction="vertical">
                    <Text strong>Compact UI mode</Text>
                    <Switch
                            checkedChildren={<CheckOutlined />}
                            unCheckedChildren={<CloseOutlined />}
                            onChange={updateCompactUIToggle}
                            checked={compactUi}
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
            <div style={{marginBottom: '8px'}}>{loggedInUser.privileges.map(privilege => <BlueTag label={privilege}/>)}</div>
            </Card>
            </div>
        </div>
    )
}

export default MyPage;