/* eslint-disable @typescript-eslint/no-explicit-any */
import { useSelector } from 'react-redux';
import { Table, Button, Popconfirm, Typography } from 'antd';
import { State } from '../../interfaces/state';
import { QuestionCircleOutlined, DeleteOutlined, ZoomInOutlined } from '@ant-design/icons';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import { Product } from '../../interfaces/product';

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
    const products = useSelector((state : State) => state.application.products);
    const userPrivileges = useSelector((state : State) => state.user.privileges);

    const productsData: Array<any> = products.map((product : Product) => {
        return {                    
            name: <Link href={ `/product/${product.id}`}>{product.name}</Link>,
            description: product.description,
            category: product.category,
            price: product.price,
            tax: `${product.tax_percentage}%`,
            operations: (
                <div style={{display: 'flex', justifyContent: 'flex-end'}}>
                    <Link style={{padding: '5px'}} href={ `/product/${product.id}`}><ZoomInOutlined /></Link>
                    {hasPrivilege(userPrivileges, PRIVILEGES.product_sudo) &&
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

    return  <><Table size="small" bordered columns={columns} dataSource={productsData} /></>
}

export default Products;