import { useSelector } from 'react-redux';
import { State } from '../types/state';
import { Typography, Tag } from 'antd';

const { Text, Title } = Typography

const MyPage: React.FC = () => {
    const user = useSelector((state : State) => state.user)

    if (!user) {
        return <Title level={1}>Something went wrong with fetching your details</Title>;
    }

    return (
        <div style={{padding: '16px 20px'}}>
            <Title level={3} style={{}}>{user.firstName} {user.lastName}</Title>
            <Title level={5}>Email</Title>
            <Text>{user.email}</Text>
            <Title level={5}>ID</Title>
            <Text>{user.id}</Text>
            <Title level={5}>Privileges</Title>
            <Text>{user.privileges.map(privilege => <Tag color="blue">{privilege}</Tag>)}</Text>
            <Title level={5}>Projects</Title>
            <Text>{user.projects.map(project => ` ${project} `)}</Text>
        </div>
    )
}

export default MyPage;