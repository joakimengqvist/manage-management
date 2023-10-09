import { Tag, Typography } from 'antd';

const { Text } = Typography;

const statusColors = {
    completed: 'green',
    ongoing: 'blue',
    idle: 'blue',
    pending: 'orange',
    cancelled: 'red',
    overdue: 'purple',
    partially_paid: 'purple',
    payment_in_progress: 'blue',
    refunded: 'gold',
    disputed: 'red',
    write_off: 'volcano',
    on_hold: 'purple',
    scheduled: 'purple',
    under_review: 'blue',
    pending_approval: 'blue',
    payment_failed: 'red',
    closed: 'default',
};

export type PaymentStatusTypes = keyof typeof statusColors;

export const ExpenseAndIncomeStatus = ({status} : {status : PaymentStatusTypes}) => {
    const color = statusColors[status] || 'default';
    return (
        <Tag color={color}>
            <Text style={{color: color}}>{status}</Text>
        </Tag>
    )
}