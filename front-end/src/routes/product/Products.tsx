import { Row, Col } from 'antd';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import CreateProduct from '../../components/products/CreateProduct';
import Products from '../../components/products/Products';
import { useGetLoggedInUserPrivileges } from '../../hooks/useGetLoggedInUserPrivileges';

const ProjectDetails = () => {
    const userPrivileges = useGetLoggedInUserPrivileges();
    return (
        <Row>
            <Col span={16}>
            <div style={{paddingRight: '8px'}}>
                {hasPrivilege(userPrivileges, PRIVILEGES.product_read) && <Products />}
            </div>
            </Col>
            <Col span={8}>
                {hasPrivilege(userPrivileges, PRIVILEGES.product_write) && <CreateProduct />}
            </Col>
        </Row>
    )

}

export default ProjectDetails;