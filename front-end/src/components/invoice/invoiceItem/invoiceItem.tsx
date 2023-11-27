/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { Col, Row, Typography } from 'antd';
import { useSelector } from 'react-redux';
import { Space, Card, notification } from 'antd';
import { State } from '../../../interfaces/state';
import { getInvoiceItemById } from '../../../api/invoices/invoiceItem/getById';
import { useParams } from 'react-router-dom';
import { formatDateTimeToYYYYMMDDHHMM } from '../../../helpers/stringDateFormatting';
import { formatNumberWithSpaces } from '../../../helpers/stringFormatting';

const { Title, Text, Link } = Typography;

const InvoiceItem = () => {
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);
    const products = useSelector((state : State) => state.application.products);
    const users = useSelector((state : State) => state.application.users);
    const [invoiceItem, setInvoiceItem] = useState<any>({});
    const { id } =  useParams(); 
    const invoiceItemId = id || '';
    
    const getProductName = (id : string) => products.find(product => product.id === id)?.name || 'Unknown';
    const getUserName = (userId : string) => users.find(user => user.id === userId)?.first_name || 'Unknown';

    useEffect(() => {
        getInvoiceItemById(userId,invoiceItemId).then(response => {
            if (response?.error || !response?.data) {
                api.error({
                    message: `Create project invoiceItem failed`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                    });
                return
            }
            setInvoiceItem(response.data);
        })
        .catch(error => {
            api.error({
                message: `Error creating project invoiceItem`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    }, []);

  return (
        <Card style={{maxWidth: '1120px '}}>
            {contextHolder}
            <Title level={4}>Invoice item</Title>
            <Row>
                <Col span={8} style={{padding: '12px 12px 12px 0px'}}>
                    <Space direction="vertical" style={{width: '100%'}}>
                        <Text strong>Product</Text>
                        <Link href={`/product/${invoiceItem.product_id}`}>{getProductName(invoiceItem.product_id)}</Link>
                        <Text strong>Total price</Text>
                        {formatNumberWithSpaces(invoiceItem.total_price)}
                        <Text strong>Tax percentage</Text>
                        {invoiceItem.tax_percentage}
                        <Text strong>Total tax</Text>
                        {formatNumberWithSpaces(invoiceItem.total_tax)}
                        <Text strong>Discount percentage</Text>
                        {invoiceItem.discount_percentage}
                        <Text strong>Quantity</Text>
                        {invoiceItem.quantity}
                    </Space>
                </Col>
                <Col span={8} style={{padding: '12px 12px 12px 0px'}}>
                    <Space direction="vertical" style={{width: '100%'}}>
                        <Text strong>Created by</Text>
                        <Link href={`/user/${invoiceItem.created_by}`}>{getUserName(invoiceItem.created_by)}</Link>
                        <Text strong>Created at</Text>
                        {formatDateTimeToYYYYMMDDHHMM(invoiceItem.created_at)}
                        <Text strong>Updated by</Text>
                        <Link href={`/user/${invoiceItem.created_by}`}>{getUserName(invoiceItem.updated_by)}</Link>
                        <Text strong>Updated at</Text>
                        {formatDateTimeToYYYYMMDDHHMM(invoiceItem.updated_at)}
                    </Space>
                </Col>
            </Row>
        </Card>
  );
};

export default InvoiceItem;