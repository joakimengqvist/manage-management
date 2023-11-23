import { Row, Col } from 'antd';
import { useSelector } from 'react-redux';
import { State } from '../../types/state';
import { hasPrivilege } from '../../helpers/hasPrivileges';
import { PRIVILEGES } from '../../enums/privileges';
import CreateProduct from '../../components/products/CreateProduct';
import Products from '../../components/products/Products';

const ProjectDetails: React.FC = () => {
    const userPrivileges = useSelector((state : State) => state.user.privileges)
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