import { Tag, Typography } from 'antd';

const { Text } = Typography;

const statusColors = {
    active: 'green',
    inactive: 'blue',
    pending_approval: 'orange',
    preferred_vendor: 'blue',
    under_review: 'blue',
    on_hold: 'orange',
    pending_payment: 'red',
    contracted: 'purple',
    prospective_partner: 'orange',
    suspended: 'red',
    closed: 'default',
    payment_issue: 'red',
  };

export type ExternalCompanyStatusTypes = keyof typeof statusColors;

export const ExternalCompanyStatus = ({status} : {status : ExternalCompanyStatusTypes}) => {
    const color = statusColors[status] || 'default';
    return (
        <Tag color={color}>
            <Text style={{color: color}}>{status.replace(/_/g, ' ')}</Text>
        </Tag>
    )
}