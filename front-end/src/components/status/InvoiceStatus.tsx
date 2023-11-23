import { Badge } from 'antd';

const InvoiceStatus = ({status} : { status : string}) => {
  switch (status) {
    case 'paid':
      return <Badge status="success" text="Paid" />;
    case 'not-sent':
      return <Badge status="default" text="Not sent" />;
    case 'sent':
      return <Badge status="processing" text="Sent" />;
    case 'delayed':
      return <Badge status="warning" text="Delayed" />;
    case 'processing':
      return <Badge status="processing" text="Processing" />;
    case 'cancelled':
      return <Badge status="error" text="Cancelled" />;
    default:
      return <Badge status="default" text="Not sent" />;
  }

}

export default InvoiceStatus;