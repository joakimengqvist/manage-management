import { Tag, Typography } from 'antd';

const { Text } = Typography;

const statusColors = {
    completed: 'green',
    ongoing: 'blue',
    cancelled: 'red',
    default: 'default',
};

export type ProjectStatusTypes = keyof typeof statusColors;

export const ProjectStatus = ({status} : {status : ProjectStatusTypes }) => {
    const color = statusColors[status];
    return (
        <Tag color={color}>
            <Text style={{color: color}}>{status}</Text>
        </Tag>
    )
}
