<?php

namespace App\DataFixtures;

use App\Entity\Product;
use Doctrine\Bundle\FixturesBundle\Fixture;
use Doctrine\Persistence\ObjectManager;

class ProductFixtures extends Fixture
{
    public function load(ObjectManager $manager): void
    {
        for ($i = 0; $i < 20; $i++) {
            $prod = new Product();
            $prod->setName("raven eisbijl");
            $prod->setMpn("123123123");
            $prod->setSource("klimwinkel.nl");
            $prod->setScrapeTime(date_create_immutable());

            $manager->persist($prod);

            $manager->flush();
        }

    }
}