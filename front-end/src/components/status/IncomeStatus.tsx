import { Badge } from 'antd';

export const IncomeStatus = ({status} : {status : string}) => {
    switch (status) {
        case 'pending':
            return <Badge status="default" text="Pending" />;
        case 'cancelled':
            return <Badge status="error" text="Cancelled" />;
        case 'overdue':
            return <Badge status="warning" text="Overdue" />;
        case 'partially_paid':
            return <Badge status="warning" text="Partially paid" />;
        case 'payment_in_progress':
            return <Badge status="processing" text="Payment in progress" />;
        case 'refunded':
            return <Badge status="success" text="Refunded" />;
        case 'disputed':
            return <Badge status="error" text="Disputed" />;
        case 'on_hold':
            return <Badge status="processing" text="On hold" />;
        case 'scheduled':
            return <Badge status="processing" text="Scheduled" />;
        case 'closed':
            return <Badge status="default" text="Closed" />;
        default:
            return <Badge status="default" text="None" />;
    }
}

export default IncomeStatus;