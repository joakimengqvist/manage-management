import { Tag, Typography } from 'antd';

const { Text } = Typography;

export const RenderProjectStatus = ({status} : {status : string}) => {
    switch(status) {
        case 'completed': {
            return (
                <Tag color="green">
                    <Text style={{color: '#389e0d'}}>{status}</Text>
                </Tag>
            )
        }
        case 'ongoing': {
            return (
                <Tag color="geekblue">
                    <Text style={{color: '#0958d9'}}>{status}</Text>
                </Tag>
            )
        }
        default: {
            return (
                <Tag>
                    <Text style={{color: 'rgba(0, 0, 0, 0.88)'}}>{status}</Text>
                </Tag>
            )
        }
    }
};