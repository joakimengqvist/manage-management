import { Badge } from 'antd';

export const ExternalCompanyStatus = ({status} : {status : string}) => {
    switch (status) {
        case 'active':
            return <Badge status="success" text="Active" />;
        case 'inactive':
            return <Badge status="error" text="Inactive" />;
        case 'pending_approval':
            return <Badge status="processing" text="Pending approval" />;
        case 'under_review':
            return <Badge status="processing" text="Under review" />;
        case 'contracted':
            return <Badge status="success" text="Contracted" />;
        case 'prospective_partner':
            return <Badge status="default" text="Prospective partner" />;
        case 'suspended':
            return <Badge status="error" text="Suspended" />;
        case 'closed':
            return <Badge status="error" text="Closed" />;
        case 'payment_issue':
            return <Badge status="error" text="Payment issue" />;
        default:
            return <Badge status="default" text="Active" />;
    }
}

export default ExternalCompanyStatus;