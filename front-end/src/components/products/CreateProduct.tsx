/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react-hooks/exhaustive-deps */
import { useState } from 'react';
import { useSelector } from 'react-redux';
import { Button, Input, Space, Card, Typography, notification } from 'antd';
import { State } from '../../interfaces/state';
import { createProduct } from '../../api/products/create';

const { Title, Text } = Typography;

const CreateUser = () => {
    const [api, contextHolder] = notification.useNotification();
    const userId = useSelector((state : State) => state.user.id);

    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [category, setCategory] = useState('');
    const [price, setPrice] = useState(0.0);
    const [taxPercentage, setTaxPercentage] = useState(0.0);

    const onSubmit = () => {
        createProduct(userId, name, description, category, price, taxPercentage)
            .then(response => {
                if (response?.error) {
                    api.error({
                        message: `Created product failed`,
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
                    message: `Error creating product`,
                    description: error.toString(),
                    placement: 'bottom',
                    duration: 1.4
                });
            })
    }

  return (
        <Card>
            {contextHolder}
            <Space direction="vertical" style={{width: '100%'}}>
                <Title level={4}>Create New Product</Title>
                <Text strong>Name</Text>
                <Input 
                    placeholder="Name" 
                    value={name} 
                    onChange={event => setName(event.target.value)} 
                    onBlur={event => setName(event.target.value)}
                />
                <Text strong>Description</Text>
                <Input 
                    placeholder="Description" 
                    value={description} 
                    onChange={event => setDescription(event.target.value)} 
                    onBlur={event => setDescription(event.target.value)}
                />
                <Text strong>Category</Text>
                <Input 
                    placeholder="Category" 
                    value={category} 
                    onChange={event => setCategory(event.target.value)} 
                    onBlur={event => setCategory(event.target.value)}
                />
                <Text strong>Price</Text>
                <Input 
                    placeholder="Price" 
                    type="number"
                    value={price} 
                    onChange={event => setPrice(Number(event.target.value))} 
                    onBlur={event => setPrice(Number(event.target.value))}
                />
                <Text strong>Tax</Text>
                <Input 
                    placeholder="Tax" 
                    type="number"
                    step="1"
                    max="100"
                    min="0"
                    defaultValue={0}
                    suffix="%"
                    value={taxPercentage} 
                    onChange={event => setTaxPercentage(Number(event.target.value))} 
                    onBlur={event => setTaxPercentage(Number(event.target.value))}
                />
                <div style={{display: 'flex', justifyContent: 'space-between', gap: '16px', marginTop: '8px'}}>
                    <Button type="primary" onClick={onSubmit}>Create product</Button>
                </div>
            </Space>
        </Card>
  );
};

export default CreateUser;