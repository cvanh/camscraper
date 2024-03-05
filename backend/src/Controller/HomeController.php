<?php

namespace App\Controller;

use App\Entity\Product;
use Doctrine\ORM\EntityManagerInterface;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class HomeController extends AbstractController
{
    #[Route('/')]
    public function home(EntityManagerInterface $enityManager): Response
    {
        $products = $enityManager->getRepository(Product::class)->findAll();

        return $this->render('home.twig',
            [
                'products' => $products,
                'data' => 1,
            ]
        );
    }

    #[Route('/createProducts')]
    public function createProducts(EntityManagerInterface $enityManager): Response
    {
        $prod = new Product();
        $prod->setName("raven eisbijl");
        $prod->setMpn("123123123");
        $prod->setSource("klimwinkel.nl");
        $prod->setScrapeTime(date_create_immutable());

        $enityManager->persist($prod);
        $enityManager->flush();

        return $this->render('home.twig',
            [
                'data' => 1,
            ]
        );
    }
}