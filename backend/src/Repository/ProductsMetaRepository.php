<?php

namespace App\Repository;

use App\Entity\ProductsMeta;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends ServiceEntityRepository<ProductsMeta>
 *
 * @method ProductsMeta|null find($id, $lockMode = null, $lockVersion = null)
 * @method ProductsMeta|null findOneBy(array $criteria, array $orderBy = null)
 * @method ProductsMeta[]    findAll()
 * @method ProductsMeta[]    findBy(array $criteria, array $orderBy = null, $limit = null, $offset = null)
 */
class ProductsMetaRepository extends ServiceEntityRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, ProductsMeta::class);
    }

    //    /**
    //     * @return ProductsMeta[] Returns an array of ProductsMeta objects
    //     */
    //    public function findByExampleField($value): array
    //    {
    //        return $this->createQueryBuilder('p')
    //            ->andWhere('p.exampleField = :val')
    //            ->setParameter('val', $value)
    //            ->orderBy('p.id', 'ASC')
    //            ->setMaxResults(10)
    //            ->getQuery()
    //            ->getResult()
    //        ;
    //    }

    //    public function findOneBySomeField($value): ?ProductsMeta
    //    {
    //        return $this->createQueryBuilder('p')
    //            ->andWhere('p.exampleField = :val')
    //            ->setParameter('val', $value)
    //            ->getQuery()
    //            ->getOneOrNullResult()
    //        ;
    //    }
}
