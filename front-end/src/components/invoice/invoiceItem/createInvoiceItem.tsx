/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from 'react';
import { Col, Row, Typography } from 'antd';

import { Button, Input, Space, Card, notification, Select } from 'antd';
import { createInvoiceItem } from '../../../api/invoices/invoiceItem/create';
import { formatNumberWithSpaces } from '../../../helpers/stringFormatting';
import { useGetLoggedInUserId, useGetProducts } from '../../../hooks';

const { Title, Text } = Typography;

const CreateInvoiceItem = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUserId = useGetLoggedInUserId();
    const products = useGetProducts();
    const [productId, setProductId] = useState('');
    const [originalPrice, setOriginalPrice] = useState(0.0);
    const [actualPrice, setActualPrice] = useState(0.0);
    const [actualTax, setActualTax] = useState(0.0);
    const [originalTax, setOriginalTax] = useState(0.0);
    const [quantity, setQuantity] = useState(0);
    const [discountPercentage, setDiscountPercentage] = useState(0);
    const [discountAmount, setDiscountAmount] = useState(0.0);
    const [selectedProduct, setSelectedProduct] = useState<any>(null);

    useEffect(() => {
        if (selectedProduct) {
            const price = selectedProduct.price * quantity;
            const initialTax = (price * (1 + selectedProduct.tax_percentage / 100)) - price;
            setOriginalPrice(price)
            setOriginalTax(initialTax);

            const finalPrice = price - price * (discountPercentage / 100);
            const finalTax = (finalPrice * (1 + selectedProduct.tax_percentage / 100)) - finalPrice;
            const discountAmount = price * (discountPercentage / 100);

            setActualPrice(finalPrice);
            setActualTax(finalTax);
            setDiscountAmount(discountAmount);

        }
    }, [quantity, selectedProduct, discountPercentage]);

    const productsOptions = Object.keys(products).map(projectId => ({ 
        label: products[projectId].name, 
        value: products[projectId].id
    }));

    const onChangeProduct = (value : string) =>  {
        const product = products[value];
        if (product) {
            setSelectedProduct(product);
            setOriginalPrice(product.price * quantity);
        }
        setProductId(value);
    }

    const onSubmit = () => {
        createInvoiceItem(
            productId,
            originalPrice,
            actualPrice,
            selectedProduct.tax_percentage,
            originalTax,
            actualTax, 
            discountPercentage,
            discountAmount,
            quantity, 
            loggedInUserId,
        ).then(response => {
            if (response?.error || !response?.data) {
                api.error({
                    message: `Create project invoiceItem failed`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                    });
                return
            }
            api.info({
                message: response.message,
                placement: 'bottom',
                duration: 1.4
            });
        })
        .catch(error => {
            api.error({
                message: `Error creating project invoiceItem`,
                description: error.toString(),
                placement: 'bottom',
                duration: 1.4
            });
        })
    };

  return (
        <Card style={{maxWidth: '1120px'}}>
            {contextHolder}
            <Title level={4}>Create invoice item</Title>
                    <Space direction="vertical" style={{width: '100%', gap: '0px'}}>
                        <Text strong>Product</Text>
                        <Select
                            placeholder="Product"
                            style={{width: '100%'}}
                            onChange={onChangeProduct}
                            value={productId}
                            options={productsOptions}
                        />
                        <Text strong>Quantity</Text>
                        <Input
                            placeholder="Quantity"
                            type="number"
                            style={{maxWidth: '144px'}}
                            onChange={(event) => setQuantity(Number(event.target.value))}
                            value={quantity}
                            suffix="Nr."
                        />
                        {selectedProduct && (<>
                            <Space direction="vertical" style={{marginBottom: '0px', marginTop: '4px', gap: '0px'}}>
                                <Text strong>Product name</Text>
                                {selectedProduct.name}
                            </Space>
                            <div style={{width: '100%', display: 'flex', justifyContent: 'flex-start', alignItems: 'center', gap: '20px', marginBottom: '0px'}}>
                                <Space direction="vertical" style={{gap: '0px'}}>
                                    <Text strong>Product price</Text>
                                    {formatNumberWithSpaces(selectedProduct.price)}
                                </Space>
                                <Space direction="vertical" style={{gap: '0px'}}>
                                    <Text strong>Product tax</Text>
                                    {`${selectedProduct.tax_percentage}%`}
                                </Space>
                            </div>
                            <div style={{width: '100%', display: 'flex', justifyContent: 'flex-start', alignItems: 'center', gap: '20px', marginBottom: '4px'}}>
                                <Space direction="vertical" style={{gap: '0px'}}>
                                    <Text strong>Total price (no discount)</Text>
                                    <Text>{formatNumberWithSpaces(originalPrice)}</Text>
                                </Space>
                            </div>
                        </>)}
                        <Text strong>Discount percentage</Text>
                        <Input
                            placeholder="Discount percentage"
                            type="number"
                            style={{maxWidth: '144px'}}
                            onChange={(event) => setDiscountPercentage(Number(event.target.value))}
                            value={discountPercentage}
                            suffix="%"
                        />
                        <div style={{width: '100%', display: 'flex', justifyContent: 'flex-start', alignItems: 'center', gap: '20px', marginBottom: '12px', marginTop: '4px'}}>
                            <Space direction="vertical" style={{gap: '0px'}}>
                                <Text strong>Discount</Text>
                                <Text>{formatNumberWithSpaces(discountAmount)}</Text>
                            </Space>
                            <Space direction="vertical" style={{gap: '0px'}}>
                                <Text strong>Final price</Text>
                                <Text>{formatNumberWithSpaces(actualPrice)}</Text>
                            </Space>
                            <Space direction="vertical" style={{gap: '0px'}}>
                                <Text strong>Total tax</Text>
                                <Text>{formatNumberWithSpaces(actualTax)}</Text>
                            </Space>
                        </div>
                    </Space>
            <Row>
                <Col>
                    <div style={{display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px'}}>
                        <Button type="primary" onClick={onSubmit}>Create invoice item</Button>
                    </div>
                </Col>
            </Row>
        </Card>
  );
};

export default CreateInvoiceItem;