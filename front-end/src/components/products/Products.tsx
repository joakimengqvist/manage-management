/* eslint-disable @typescript-eslint/no-explicit-any */
import { Table, Button, Popconfirm, Typography, notification } from 'antd';
import { QuestionCircleOutlined, DeleteOutlined, ZoomInOutlined } from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { Product } from '../../interfaces/product';
import { useEffect, useState } from 'react';
import { getAllProducts } from '../../api';
import { useGetLoggedInUser } from '../../hooks';
import { formatNumberWithSpaces } from '../../helpers/stringFormatting';

const { Link } = Typography;

const columns = [
    {
        title: 'name',
        dataIndex: 'name',
        key: 'name'
    },
    {
        title: 'Description',
        dataIndex: 'description',
        key: 'description'
    },
    {
        title: 'Category',
        dataIndex: 'category',
        key: 'category'
    },
    {
        title: 'Price',
        dataIndex: 'price',
        key: 'price'
    },
    {
        title: 'Tax',
        dataIndex: 'tax',
        key: 'tax'
    },
    {
        title: '',
        dataIndex: 'operations',
        key: 'operations'
    },
]

const Products = () => {
    const [api, contextHolder] = notification.useNotification();
    const loggedInUser = useGetLoggedInUser();
    const [products, setProducts] = useState<Array<Product>>([]);

    useEffect(() => {
        getAllProducts(loggedInUser.id).then(response => {
            if (response?.error) {
                api.error({
                    message: `Error fetching privileges`,
                    description: response.message,
                    placement: 'bottom',
                    duration: 1.4
                  });
                  return
              }
              setProducts(response.data)
        })
    }, [api, loggedInUser.id])

    const productsData: Array<any> = products.map((product : Product) => {
        return {                    
            name: <Link href={ `/product/${product.id}`}>{product.name}</Link>,
            description: product.description,
            category: product.category,
            price: `${formatNumberWithSpaces(product.price)} SEK`,
            tax: `${product.tax_percentage}%`,
            operations: (
                <div style={{display: 'flex', justifyContent: 'flex-end'}}>
                    <Link style={{padding: '5px'}} href={ `/product/${product.id}`}><ZoomInOutlined /></Link>
                    {hasPrivilege(loggedInUser.privileges, PRIVILEGES.product_sudo) &&
                        <Popconfirm
                            placement="top"
                            title="Are you sure?"
                            description={`Do you want to delete product ${product.name}`}
                            onConfirm={() => alert(`Deleting ${product.name}`)}
                            icon={<QuestionCircleOutlined style={{ color: 'red' }} />}
                            okText="Yes"
                            cancelText="No"
                        >
                            <Button style={{ padding: '4px' }} danger type="link"><DeleteOutlined /></Button>
                        </Popconfirm>
                    }
                </div>
            )
        }
    });

    return  <>{contextHolder}<Table size="small" bordered columns={columns} dataSource={productsData} /></>
}

export default Products;